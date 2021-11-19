package errors

type KZerosControlPlaneStatusError string

const (
	// InvalidConfigurationKZerosControlPlaneError indicates that the KZeros control plane
	// configuration is invalid.
	InvalidConfigurationKZerosControlPlaneError KZerosControlPlaneStatusError = "InvalidConfiguration"

	// UnsupportedChangeKZerosControlPlaneError indicates that the KZeros control plane
	// spec has been updated in an unsupported way that cannot be
	// reconciled.
	UnsupportedChangeKZerosControlPlaneError KZerosControlPlaneStatusError = "UnsupportedChange"

	// CreateKZerosControlPlaneError indicates that an error was encountered
	// when trying to create the KZeros control plane.
	CreateKZerosControlPlaneError KZerosControlPlaneStatusError = "CreateError"

	// UpdateKZerosControlPlaneError indicates that an error was encountered
	// when trying to update the KZeros control plane.
	UpdateKZerosControlPlaneError KZerosControlPlaneStatusError = "UpdateError"

	// DeleteKZerosControlPlaneError indicates that an error was encountered
	// when trying to delete the KZeros control plane.
	DeleteKZerosControlPlaneError KZerosControlPlaneStatusError = "DeleteError"
)
