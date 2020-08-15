package request

type CustomClaims struct {
	AdminId          float64 `gconv:"admin_id"`
	AdminUuid        string  `gconv:"admin_uuid"`
	AdminNickname    string  `gconv:"admin_nickname"`
	Exp              float64 `gconv:"exp"`
	OrigIat          float64 `gconv:"orig_iat"`
	AdminAuthorityId string  `gconv:"admin_authority_id"`
}
