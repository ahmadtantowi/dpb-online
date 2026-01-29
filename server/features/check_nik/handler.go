package check_nik

import (
	"database/sql"
	"dpb-online/db"
	"dpb-online/server/dto"
	"log/slog"
	"net/http"
)

type CheckNIKResp struct {
	NIK         string `json:"nik"`
	NKK         string `json:"nkk"`
	Name        string `json:"name"`
	City        string `json:"city"`
	District    string `json:"district"`
	Subdistrict string `json:"subdistrict"`
	NRT         string `json:"nrt"`
	NRW         string `json:"nrw"`
	TPSNo       string `json:"tps_no"`
	TPSAddress  string `json:"tps_address"`
}

// @Summary Check NIK
// @Description Lookup pemilih (voter) data by NIK.
// @Tags Voter
// @Produce application/json
// @Param nik path string true "NIK"
// @Success 200 {object} dto.Response[CheckNIKResp]
// @Failure 400 {object} dto.Response[string]
// @Failure 404 {object} dto.Response[string]
// @Failure 500 {object} dto.Response[string]
// @Router /{nik} [get]
func CheckNIKHandler(logger *slog.Logger, db *db.Database) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nik := r.PathValue("nik")
		if nik == "" {
			dto.JsonErr(w, http.StatusBadRequest, "NIK is required")
			return
		}

		dpb, err := getPemilihByNIK(r.Context(), db, nik)
		if err != nil {
			if err == sql.ErrNoRows {
				dto.JsonErr(w, http.StatusNotFound, "NIK not found")
			} else {
				dto.JsonErr(w, http.StatusInternalServerError, err.Error())
			}
			return
		}

		resp := CheckNIKResp{
			NIK:         dpb.nik,
			NKK:         dpb.nik,
			Name:        dpb.name,
			City:        dpb.city,
			District:    dpb.district,
			Subdistrict: dpb.subdistrict,
			NRT:         dpb.nrt,
			NRW:         dpb.nrw,
			TPSNo:       dpb.tpsNo,
			TPSAddress:  dpb.tpsAddress,
		}

		dto.JsonOK(w, http.StatusOK, resp)
	})
}
