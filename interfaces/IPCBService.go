package interfaces

type IPCBService interface {
	RETRCOMACCTDTLSMW(acctno string) (string, error)
}
