package restore

import (
	"github.com/caos/zitadel/operator/database/kinds/backups/core"
	"time"

	"github.com/caos/zitadel/operator"

	"github.com/caos/orbos/mntr"
	"github.com/caos/orbos/pkg/kubernetes"
	"github.com/caos/orbos/pkg/kubernetes/resources/job"
	"github.com/caos/orbos/pkg/labels"
	corev1 "k8s.io/api/core/v1"
)

const (
	Instant            = "restore"
	defaultMode        = int32(256)
	certPath           = "/cockroach/cockroach-certs"
	secretPath         = "/secrets/sa.json"
	internalSecretName = "client-certs"
	rootSecretName     = "cockroachdb.client.root"
	timeout            = 45 * time.Minute
	saJsonBase64Env    = "SAJSON"
)

func AdaptFunc(
	monitor mntr.Monitor,
	backupName string,
	namespace string,
	componentLabels *labels.Component,
	bucketName string,
	timestamp string,
	nodeselector map[string]string,
	tolerations []corev1.Toleration,
	checkDBReady operator.EnsureFunc,
	secretKey string,
	dbURL string,
	dbPort int32,
	image string,
) (
	queryFunc operator.QueryFunc,
	destroyFunc operator.DestroyFunc,
	err error,
) {

	jobName := core.GetRestoreJobName(backupName)
	command := getCommand(
		timestamp,
		bucketName,
		backupName,
		certPath,
		secretPath,
		dbURL,
		dbPort,
	)

	jobdef := getJob(
		namespace,
		labels.MustForName(componentLabels, jobName),
		nodeselector,
		tolerations,
		core.GetSecretName(backupName),
		secretKey,
		command,
		image,
	)

	destroyJ, err := job.AdaptFuncToDestroy(jobName, namespace)
	if err != nil {
		return nil, nil, err
	}

	destroyers := []operator.DestroyFunc{
		operator.ResourceDestroyToZitadelDestroy(destroyJ),
	}

	queryJ, err := job.AdaptFuncToEnsure(jobdef)
	if err != nil {
		return nil, nil, err
	}

	queriers := []operator.QueryFunc{
		operator.EnsureFuncToQueryFunc(checkDBReady),
		operator.ResourceQueryToZitadelQuery(queryJ),
	}

	return func(k8sClient kubernetes.ClientInt, queried map[string]interface{}) (operator.EnsureFunc, error) {
			return operator.QueriersToEnsureFunc(monitor, false, queriers, k8sClient, queried)
		},
		operator.DestroyersToDestroyFunc(monitor, destroyers),

		nil
}
