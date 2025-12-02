package check_nik

import (
	"context"
	"dpb-online/db"
)

type pemilih struct {
	nkk, nik              string
	name                  string
	city                  string // kab/kota
	district, subdistrict string // kecamatan, desa/kelurahan
	address               string
	nrt, nrw              string
	tpsNo, tpsAddress     string
}

func getPemilihByNIK(ctx context.Context, db *db.Database, nik string) (*pemilih, error) {
	const q = `
	select p.nkk ,p.nik ,p.nama ,kab.nama kab ,kec.nama kec ,kel.nama kel ,p.alamat ,p.rt ,p.rw ,t.tps_no ,t.alamat_tps
	from pemilih p
	left join wilayah kab on kab.wilayah_id = p.kab_id
	left join wilayah kec on kec.wilayah_id = p.kec_id
	left join wilayah kel on kel.wilayah_id = p.kel_id
	left join tps t on t.tps_id = p.tps_id
	where 1=1
		and p.status not in ('tms', 'delete')
		and nik = ?`

	var pemilih pemilih
	err := db.DB().QueryRowContext(ctx, q, nik).Scan(
		&pemilih.nkk,
		&pemilih.nik,
		&pemilih.name,
		&pemilih.city,
		&pemilih.district,
		&pemilih.subdistrict,
		&pemilih.address,
		&pemilih.nrt,
		&pemilih.nrw,
		&pemilih.tpsNo,
		&pemilih.tpsAddress,
	)
	if err != nil {
		return nil, err
	}

	return &pemilih, nil
}
