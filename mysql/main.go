package mysql

import (
	"github.com/albatiqy/gopoh"
	"github.com/albatiqy/gopoh/provider/dblib"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var (
	AttrsMap = dblib.AttrsMap{
		"ptk_id":                       {Col: "ptk_id", Type: (*dblib.NullString)(nil)},
		"nama":                         {Col: "nama", Type: (*string)(nil)},
		"nip":                          {Col: "nip", Type: (*dblib.NullString)(nil)},
		"jenis_kelamin":                {Col: "jenis_kelamin", Type: (*string)(nil)},
		"tempat_lahir":                 {Col: "tempat_lahir", Type: (*string)(nil)},
		"tanggal_lahir":                {Col: "tanggal_lahir", Type: (*time.Time)(nil)},
		"nik":                          {Col: "nik", Type: (*string)(nil)},
		"no_kk":                        {Col: "no_kk", Type: (*dblib.NullString)(nil)},
		"niy_nigk":                     {Col: "niy_nigk", Type: (*dblib.NullString)(nil)},
		"nuptk":                        {Col: "nuptk", Type: (*dblib.NullString)(nil)},
		"nuks":                         {Col: "nuks", Type: (*dblib.NullString)(nil)},
		"status_kepegawaian_id":        {Col: "status_kepegawaian_id", Type: (*int16)(nil)},
		"jenis_ptk_id":                 {Col: "jenis_ptk_id", Type: (*float32)(nil)},
		"pengawas_bidang_studi_id":     {Col: "pengawas_bidang_studi_id", Type: (*int32)(nil)},
		"agama_id":                     {Col: "agama_id", Type: (*int16)(nil)},
		"kewarganegaraan":              {Col: "kewarganegaraan", Type: (*string)(nil)},
		"alamat_jalan":                 {Col: "alamat_jalan", Type: (*string)(nil)},
		"rt":                           {Col: "rt", Type: (*float32)(nil)},
		"rw":                           {Col: "rw", Type: (*float32)(nil)},
		"nama_dusun":                   {Col: "nama_dusun", Type: (*dblib.NullString)(nil)},
		"desa_kelurahan":               {Col: "desa_kelurahan", Type: (*string)(nil)},
		"kode_wilayah":                 {Col: "kode_wilayah", Type: (*string)(nil)},
		"kode_pos":                     {Col: "kode_pos", Type: (*dblib.NullString)(nil)},
		"lintang":                      {Col: "lintang", Type: (*float32)(nil)},
		"bujur":                        {Col: "bujur", Type: (*float32)(nil)},
		"no_telepon_rumah":             {Col: "no_telepon_rumah", Type: (*dblib.NullString)(nil)},
		"no_hp":                        {Col: "no_hp", Type: (*dblib.NullString)(nil)},
		"email":                        {Col: "email", Type: (*dblib.NullString)(nil)},
		"status_keaktifan_id":          {Col: "status_keaktifan_id", Type: (*float32)(nil)},
		"sk_cpns":                      {Col: "sk_cpns", Type: (*dblib.NullString)(nil)},
		"tgl_cpns":                     {Col: "tgl_cpns", Type: (*dblib.NullTime)(nil)},
		"sk_pengangkatan":              {Col: "sk_pengangkatan", Type: (*dblib.NullString)(nil)},
		"tmt_pengangkatan":             {Col: "tmt_pengangkatan", Type: (*dblib.NullTime)(nil)},
		"lembaga_pengangkat_id":        {Col: "lembaga_pengangkat_id", Type: (*float32)(nil)},
		"pangkat_golongan_id":          {Col: "pangkat_golongan_id", Type: (*float32)(nil)},
		"keahlian_laboratorium_id":     {Col: "keahlian_laboratorium_id", Type: (*int16)(nil)},
		"sumber_gaji_id":               {Col: "sumber_gaji_id", Type: (*float32)(nil)},
		"nama_ibu_kandung":             {Col: "nama_ibu_kandung", Type: (*string)(nil)},
		"status_perkawinan":            {Col: "status_perkawinan", Type: (*float32)(nil)},
		"nama_suami_istri":             {Col: "nama_suami_istri", Type: (*dblib.NullString)(nil)},
		"nip_suami_istri":              {Col: "nip_suami_istri", Type: (*dblib.NullString)(nil)},
		"pekerjaan_suami_istri":        {Col: "pekerjaan_suami_istri", Type: (*int32)(nil)},
		"tmt_pns":                      {Col: "tmt_pns", Type: (*dblib.NullTime)(nil)},
		"sudah_lisensi_kepala_sekolah": {Col: "sudah_lisensi_kepala_sekolah", Type: (*float32)(nil)},
		"jumlah_sekolah_binaan":        {Col: "jumlah_sekolah_binaan", Type: (*int16)(nil)},
		"pernah_diklat_kepengawasan":   {Col: "pernah_diklat_kepengawasan", Type: (*float32)(nil)},
		"nm_wp":                        {Col: "nm_wp", Type: (*dblib.NullString)(nil)},
		"status_data":                  {Col: "status_data", Type: (*int32)(nil)},
		"karpeg":                       {Col: "karpeg", Type: (*dblib.NullString)(nil)},
		"karpas":                       {Col: "karpas", Type: (*dblib.NullString)(nil)},
		"mampu_handle_kk":              {Col: "mampu_handle_kk", Type: (*int32)(nil)},
		"keahlian_braille":             {Col: "keahlian_braille", Type: (*float32)(nil)},
		"keahlian_bhs_isyarat":         {Col: "keahlian_bhs_isyarat", Type: (*float32)(nil)},
		"npwp":                         {Col: "npwp", Type: (*dblib.NullString)(nil)},
		"id_bank":                      {Col: "id_bank", Type: (*dblib.NullString)(nil)},
		"rekening_bank":                {Col: "rekening_bank", Type: (*dblib.NullString)(nil)},
		"rekening_atas_nama":           {Col: "rekening_atas_nama", Type: (*dblib.NullString)(nil)},
		"blob_id":                      {Col: "blob_id", Type: (*dblib.NullString)(nil)},
		"create_date":                  {Col: "create_date", Type: (*time.Time)(nil)},
		"Last_update":                  {Col: "Last_update", Type: (*time.Time)(nil)},
		"Soft_delete":                  {Col: "Soft_delete", Type: (*float32)(nil)},
		"last_sync":                    {Col: "last_sync", Type: (*time.Time)(nil)},
		"Updater_ID":                   {Col: "Updater_ID", Type: (*dblib.NullString)(nil)},
	}
	clientDbQuery dblib.SQLDbQuery
)

func GetDbQuery() dblib.SQLDbQuery {
	if clientDbQuery == nil {
		mysql := dblib.GetMySql()
		clientDbQuery = mysql.NewSQLDbQuery("ptk", "ptk_id", AttrsMap, gopoh.Attrs{
			"client_id",
			"secret",
			"email",
			"user_id",
			"aktif",
		}, false)
	}
	return clientDbQuery
}
