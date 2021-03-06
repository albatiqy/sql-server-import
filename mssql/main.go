package mssql

import (
	"time"

	"github.com/albatiqy/gopoh"
	"github.com/albatiqy/gopoh/provider/dblib"
)

var (
	clientDbQuery dblib.SQLDbQuery
)

func GetDbQuery() dblib.SQLDbQuery {
	if clientDbQuery == nil {
		mysql := dblib.GetMsSql()
		clientDbQuery = mysql.NewSQLDbQuery("dbo.ptk", "ptk_id", dblib.AttrsMap{
			"nm_wp":	{Col: "nm_wp", Type: (*dblib.NullString)(nil)},
			"niy_nigk":	{Col: "niy_nigk", Type: (*dblib.NullString)(nil)},
			"pengawas_bidang_studi_id":	{Col: "pengawas_bidang_studi_id", Type: (*dblib.NullInt32)(nil)},
			"email":	{Col: "email", Type: (*dblib.NullString)(nil)},
			"sumber_gaji_id":	{Col: "sumber_gaji_id", Type: (*float32)(nil)},
			"kewarganegaraan":	{Col: "kewarganegaraan", Type: (*string)(nil)},
			"bujur":	{Col: "bujur", Type: (*dblib.NullFloat64)(nil)},
			"status_perkawinan":	{Col: "status_perkawinan", Type: (*float32)(nil)},
			"nama_ibu_kandung":	{Col: "nama_ibu_kandung", Type: (*string)(nil)},
			"nama_suami_istri":	{Col: "nama_suami_istri", Type: (*dblib.NullString)(nil)},
			"sudah_lisensi_kepala_sekolah":	{Col: "sudah_lisensi_kepala_sekolah", Type: (*float32)(nil)},
			"jumlah_sekolah_binaan":	{Col: "jumlah_sekolah_binaan", Type: (*dblib.NullInt32)(nil)},
			"nip":	{Col: "nip", Type: (*dblib.NullString)(nil)},
			"tempat_lahir":	{Col: "tempat_lahir", Type: (*string)(nil)},
			"no_hp":	{Col: "no_hp", Type: (*dblib.NullString)(nil)},
			"tgl_cpns":	{Col: "tgl_cpns", Type: (*dblib.NullTime)(nil)},
			"nama":	{Col: "nama", Type: (*string)(nil)},
			"nuptk":	{Col: "nuptk", Type: (*dblib.NullString)(nil)},
			"pangkat_golongan_id":	{Col: "pangkat_golongan_id", Type: (*dblib.NullFloat64)(nil)},
			"nip_suami_istri":	{Col: "nip_suami_istri", Type: (*dblib.NullString)(nil)},
			"karpas":	{Col: "karpas", Type: (*dblib.NullString)(nil)},
			"ptk_id":	{Col: "ptk_id", Type: (*dblib.MsSqlUniqueIdentifier)(nil)},
			"jenis_kelamin":	{Col: "jenis_kelamin", Type: (*string)(nil)},
			"no_kk":	{Col: "no_kk", Type: (*dblib.NullString)(nil)},
			"alamat_jalan":	{Col: "alamat_jalan", Type: (*string)(nil)},
			"nuks":	{Col: "nuks", Type: (*dblib.NullString)(nil)},
			"agama_id":	{Col: "agama_id", Type: (*int16)(nil)},
			"nama_dusun":	{Col: "nama_dusun", Type: (*dblib.NullString)(nil)},
			"create_date":	{Col: "create_date", Type: (*time.Time)(nil)},
			"keahlian_braille":	{Col: "keahlian_braille", Type: (*dblib.NullFloat64)(nil)},
			"keahlian_bhs_isyarat":	{Col: "keahlian_bhs_isyarat", Type: (*dblib.NullFloat64)(nil)},
			"rekening_bank":	{Col: "rekening_bank", Type: (*dblib.NullString)(nil)},
			"tanggal_lahir":	{Col: "tanggal_lahir", Type: (*time.Time)(nil)},
			"no_telepon_rumah":	{Col: "no_telepon_rumah", Type: (*dblib.NullString)(nil)},
			"tmt_pengangkatan":	{Col: "tmt_pengangkatan", Type: (*dblib.NullTime)(nil)},
			"status_data":	{Col: "status_data", Type: (*dblib.NullInt32)(nil)},
			"nik":	{Col: "nik", Type: (*string)(nil)},
			"rw":	{Col: "rw", Type: (*dblib.NullFloat64)(nil)},
			"sk_pengangkatan":	{Col: "sk_pengangkatan", Type: (*dblib.NullString)(nil)},
			"tmt_pns":	{Col: "tmt_pns", Type: (*dblib.NullTime)(nil)},
			"kode_wilayah":	{Col: "kode_wilayah", Type: (*string)(nil)},
			"Soft_delete":	{Col: "Soft_delete", Type: (*float32)(nil)},
			"npwp":	{Col: "npwp", Type: (*dblib.NullString)(nil)},
			"last_sync":	{Col: "last_sync", Type: (*time.Time)(nil)},
			"status_kepegawaian_id":	{Col: "status_kepegawaian_id", Type: (*int16)(nil)},
			"lintang":	{Col: "lintang", Type: (*dblib.NullFloat64)(nil)},
			"status_keaktifan_id":	{Col: "status_keaktifan_id", Type: (*float32)(nil)},
			"mampu_handle_kk":	{Col: "mampu_handle_kk", Type: (*int32)(nil)},
			"Updater_ID":	{Col: "Updater_ID", Type: (*dblib.MsSqlUniqueIdentifier)(nil)},
			"sk_cpns":	{Col: "sk_cpns", Type: (*dblib.NullString)(nil)},
			"keahlian_laboratorium_id":	{Col: "keahlian_laboratorium_id", Type: (*dblib.NullInt32)(nil)},
			"pekerjaan_suami_istri":	{Col: "pekerjaan_suami_istri", Type: (*int32)(nil)},
			"pernah_diklat_kepengawasan":	{Col: "pernah_diklat_kepengawasan", Type: (*float32)(nil)},
			"kode_pos":	{Col: "kode_pos", Type: (*dblib.NullString)(nil)},
			"lembaga_pengangkat_id":	{Col: "lembaga_pengangkat_id", Type: (*float32)(nil)},
			"karpeg":	{Col: "karpeg", Type: (*dblib.NullString)(nil)},
			"Last_update":	{Col: "Last_update", Type: (*time.Time)(nil)},
			"rt":	{Col: "rt", Type: (*dblib.NullFloat64)(nil)},
			"desa_kelurahan":	{Col: "desa_kelurahan", Type: (*string)(nil)},
			"id_bank":	{Col: "id_bank", Type: (*dblib.NullString)(nil)},
			"rekening_atas_nama":	{Col: "rekening_atas_nama", Type: (*dblib.NullString)(nil)},
			"jenis_ptk_id":	{Col: "jenis_ptk_id", Type: (*float32)(nil)},
			"blob_id":	{Col: "blob_id", Type: (*dblib.MsSqlNullUniqueIdentifier)(nil)},
		}, gopoh.Attrs{
			"ptk_id",
			"nama",
			"nip",
			"jenis_kelamin",
			"tempat_lahir",
			"tanggal_lahir",
			"nik",
			"no_kk",
			"niy_nigk",
			"nuptk",
			"nuks",
			"status_kepegawaian_id",
			"jenis_ptk_id",
			"pengawas_bidang_studi_id",
			"agama_id",
			"kewarganegaraan",
			"alamat_jalan",
			"rt",
			"rw",
			"nama_dusun",
			"desa_kelurahan",
			"kode_wilayah",
			"kode_pos",
			"lintang",
			"bujur",
			"no_telepon_rumah",
			"no_hp",
			"email",
			"status_keaktifan_id",
			"sk_cpns",
			"tgl_cpns",
			"sk_pengangkatan",
			"tmt_pengangkatan",
			"lembaga_pengangkat_id",
			"pangkat_golongan_id",
			"keahlian_laboratorium_id",
			"sumber_gaji_id",
			"nama_ibu_kandung",
			"status_perkawinan",
			"nama_suami_istri",
			"nip_suami_istri",
			"pekerjaan_suami_istri",
			"tmt_pns",
			"sudah_lisensi_kepala_sekolah",
			"jumlah_sekolah_binaan",
			"pernah_diklat_kepengawasan",
			"nm_wp",
			"status_data",
			"karpeg",
			"karpas",
			"mampu_handle_kk",
			"keahlian_braille",
			"keahlian_bhs_isyarat",
			"npwp",
			"id_bank",
			"rekening_bank",
			"rekening_atas_nama",
			"blob_id",
			"create_date",
			"Last_update",
			"Soft_delete",
			"last_sync",
			"Updater_ID",
		}, false)
	}

	return clientDbQuery
}
