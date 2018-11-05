package transaction

type TxnInputAttribute struct {
	Name          string
	DataType      string
	IsAssetRef    bool
	IsArray       bool
	AssetName     string
	Identifiers   string
	IsParticipant bool
}

type TxnACL struct {
	AuthorizedParty []string
	Conditions      map[string]string
}
