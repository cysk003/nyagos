pushd "%~dp0"
go run importconst_run.go -p dos ^
	RESOURCE_CONNECTED ^
	RESOURCE_CONTEXT ^
	RESOURCE_GLOBALNET ^
	RESOURCE_REMEMBERED ^
	RESOURCETYPE_ANY ^
	RESOURCETYPE_DISK ^
	RESOURCETYPE_PRINT ^
	RESOURCEDISPLAYTYPE_NETWORK ^
	RESOURCEDISPLAYTYPE_SERVER ^
	RESOURCEUSAGE_CONNECTABLE ^
	RESOURCEUSAGE_CONTAINER ^
	RESOURCEUSAGE_ATTACHED ^
	RESOURCEUSAGE_ALL ^
	ERROR_NO_MORE_ITEMS ^
	CONNECT_UPDATE_PROFILE ^
	S_OK
popd
