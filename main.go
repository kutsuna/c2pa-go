package c2pa

// #cgo CFLAGS: -I/path/to/c2pa-c/include
// #cgo LDFLAGS: -L/home/kutsuna/libs/c2pa-c/target/release -lc2pa
// #include <c2pa.h>
import "C"

import (
	"unsafe"
)

type SignerInfo struct {
    Alg         string
    SignCert    string
    PrivateKey  string
    TaURL       string
}

type Manifest struct {
    ClaimGenerator      string          `json:"claim_generator"`
    ClaimGeneratorInfo  []GeneratorInfo `json:"claim_generator_info"`
    Assertions          []Assertion     `json:"assertions"`
}

type GeneratorInfo struct {
    Name    string `json:"name"`
    Version string `json:"version"`
}

type Assertion struct {
    Label string            `json:"label"`
    Data  AssertionData     `json:"data"`
}

type AssertionData struct {
    Entries map[string]Entry `json:"entries"`
}

type Entry struct {
    Use string `json:"use"`
}

func ReadFile(path, dataDir string) string {
    cPath := C.CString(path)
    cDataDir := C.CString(dataDir)
    defer C.free(unsafe.Pointer(cPath))
    defer C.free(unsafe.Pointer(cDataDir))

    result := C.c2pa_read_file(cPath, cDataDir)

    return C.GoString(result)
}

func convertToCSignerInfo(signerInfo SignerInfo) *C.C2paSignerInfo {
    c2paSignerInfo := &C.C2paSignerInfo{
        alg:         C.CString(signerInfo.Alg),
        sign_cert:   C.CString(signerInfo.SignCert),
        private_key: C.CString(signerInfo.PrivateKey),
        ta_url:      C.CString(signerInfo.TaURL),
    }

    return c2paSignerInfo
}

func freeCSignerInfo(cInfo *C.C2paSignerInfo) {
    C.free(unsafe.Pointer(cInfo.alg))
    C.free(unsafe.Pointer(cInfo.sign_cert))
    C.free(unsafe.Pointer(cInfo.private_key))
    C.free(unsafe.Pointer(cInfo.ta_url))
}


func SignFile(sourcePath, destPath, manifest string, signerInfo SignerInfo, dataDir string) string {
    cSourcePath := C.CString(sourcePath)
    cDestPath := C.CString(destPath)
    cManifest := C.CString(manifest)
    cDataDir := C.CString(dataDir)
    defer C.free(unsafe.Pointer(cSourcePath))
    defer C.free(unsafe.Pointer(cDestPath))
    defer C.free(unsafe.Pointer(cManifest))
    defer C.free(unsafe.Pointer(cDataDir))

	cSignerInfo := convertToCSignerInfo(signerInfo)
	defer freeCSignerInfo(cSignerInfo)

    result := C.c2pa_sign_file(cSourcePath, cDestPath, cManifest, cSignerInfo, cDataDir)
    return C.GoString(result)
}
