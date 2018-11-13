// 由GOVCL UI设计器自动生成，不要编辑
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    StatusBar1   *vcl.TStatusBar
    Panel1       *vcl.TPanel
    PnlWebview   *vcl.TPanel
    Panel2       *vcl.TPanel
    Button1      *vcl.TButton
    Panel3       *vcl.TPanel
    EdtURL       *vcl.TEdit
    Label1       *vcl.TLabel
    Panel4       *vcl.TPanel
    BtnGoBack    *vcl.TSpeedButton
    BtnGoForward *vcl.TSpeedButton
    BtnRefresh   *vcl.TSpeedButton

    //::private::
    TMainFormFields
}

var MainForm *TMainForm




// 以字节形式加载
// vcl.Application.CreateForm(mainFormBytes, &MainForm)

var (
    mainFormBytes = []byte {
        0x47, 0x4F, 0x56, 0x43, 0x4C, 0x46, 0x4F, 0x52, 0x4D, 0x5A, 0x01, 0x00, 
        0x93, 0x48, 0x06, 0xDC, 0xA8, 0x61, 0xEB, 0x1B, 0x91, 0x7E, 0xE0, 0x35, 
        0xEB, 0x5F, 0x31, 0x88, 0xC6, 0xAB, 0x97, 0xDC, 0x98, 0x8E, 0xE2, 0xFE, 
        0x2C, 0x95, 0x7E, 0xF7, 0xB5, 0x81, 0x3F, 0x88, 0x82, 0x45, 0x41, 0x6E, 
        0xDD, 0x52, 0x4D, 0x1A, 0x0B, 0x33, 0x52, 0x5B, 0x7E, 0xA8, 0x55, 0x9E, 
        0x10, 0x06, 0xEC, 0x5B, 0xB4, 0x1C, 0x0B, 0x50, 0x0B, 0xD7, 0xAF, 0x96, 
        0x28, 0xC9, 0xCD, 0xFF, 0xA3, 0x98, 0xFF, 0x34, 0x97, 0xAA, 0x9A, 0x35, 
        0x94, 0x40, 0x2F, 0xA2, 0xC7, 0xC8, 0x68, 0xB5, 0xB9, 0xFF, 0x05, 0x90, 
        0x70, 0x75, 0xC5, 0x20, 0x57, 0x5F, 0xE2, 0xA1, 0x87, 0x8B, 0x18, 0x07, 
        0x71, 0x81, 0xF0, 0xB6, 0xB7, 0xBD, 0xBD, 0xD3, 0x9E, 0x66, 0x63, 0x07, 
        0x40, 0x86, 0xEE, 0x1C, 0xB1, 0xB8, 0x9B, 0xC9, 0xC6, 0x4F, 0x7B, 0x15, 
        0x11, 0xCF, 0xD2, 0xD8, 0x31, 0x71, 0x26, 0xC4, 0x8F, 0xC3, 0xE4, 0xBE, 
        0x9F, 0xD0, 0xC7, 0xA2, 0x68, 0x7D, 0xC6, 0x69, 0xA6, 0xE8, 0xA7, 0x99, 
        0x33, 0x37, 0x08, 0xD0, 0x63, 0x52, 0x96, 0x25, 0x0F, 0x81, 0x96, 0x0C, 
        0x14, 0x1F, 0xDA, 0x69, 0x3E, 0xFD, 0x3E, 0x8A, 0xBF, 0xFD, 0xB5, 0xF7, 
        0xC4, 0x73, 0xF3, 0x79, 0x10, 0x0A, 0xBA, 0x6A, 0x32, 0x7C, 0xD0, 0x39, 
        0xEE, 0xDA, 0xD5, 0x39, 0x96, 0x83, 0xF6, 0xAE, 0xBC, 0x68, 0xD1, 0x0D, 
        0x09, 0xF6, 0xAC, 0x88, 0x4B, 0x96, 0x61, 0xF4, 0x5F, 0xF2, 0x14, 0xA7, 
        0xDA, 0x58, 0xC8, 0x0F, 0x3F, 0xBE, 0xC4, 0x65, 0xC4, 0xE0, 0x7E, 0xD7, 
        0x35, 0xF5, 0x58, 0x5B, 0xD6, 0xA1, 0xD0, 0xB7, 0xF4, 0x24, 0xCF, 0xFC, 
        0xB6, 0x5B, 0xD9, 0xF4, 0xE5, 0x0C, 0xF5, 0x94, 0x32, 0xBD, 0x46, 0xAA, 
        0x99, 0x0D, 0x93, 0x81, 0x80, 0xD5, 0xC4, 0xCB, 0x6B, 0xA4, 0x07, 0xE6, 
        0x99, 0x42, 0x5C, 0x8D, 0x09, 0x6C, 0x72, 0x73, 0x4F, 0xC1, 0x5C, 0xC1, 
        0xA5, 0x98, 0x8E, 0xF7, 0x92, 0x21, 0xC9, 0xD1, 0x30, 0x13, 0x4B, 0xE3, 
        0x1F, 0x90, 0x89, 0x64, 0x54, 0x9C, 0x28, 0x4A, 0x37, 0x5C, 0x41, 0x83, 
        0x19, 0xB8, 0x84, 0x76, 0xA2, 0x4D, 0x91, 0x03, 0xED, 0xAE, 0xB7, 0x84, 
        0xAA, 0x32, 0x93, 0xB6, 0x21, 0x1C, 0xA4, 0x87, 0x78, 0x98, 0x6B, 0xE8, 
        0xDB, 0x67, 0x6C, 0xFF, 0x00, 0xBF, 0x6A, 0x13, 0xB6, 0x07, 0x08, 0x80, 
        0x84, 0xC2, 0x3A, 0x93, 0xDC, 0xB7, 0xA9, 0x46, 0xE9, 0x95, 0x2A, 0x29, 
        0xAE, 0xDC, 0x13, 0x26, 0xD2, 0x55, 0x8A, 0x3B, 0xF0, 0xD4, 0xC9, 0x2D, 
        0x51, 0xC7, 0x19, 0xAE, 0xCE, 0xE9, 0x5C, 0x98, 0x87, 0x1B, 0x3D, 0x42, 
        0x82, 0xCD, 0x0D, 0xF9, 0x61, 0x58, 0xFE, 0x1D, 0xD9, 0x1A, 0xCF, 0xC1, 
        0x0F, 0xE4, 0x6F, 0x6E, 0x0B, 0x92, 0x3D, 0xD8, 0x78, 0xD3, 0x81, 0x7A, 
        0xA7, 0xD7, 0x42, 0x1E, 0x6F, 0xCE, 0x14, 0xB8, 0xB4, 0xF4, 0x34, 0xF1, 
        0x00, 0x19, 0xEE, 0x53, 0x07, 0xB9, 0xA0, 0x47, 0x68, 0x1B, 0x3A, 0xA2, 
        0xCD, 0x90, 0x4D, 0xAC, 0x60, 0x90, 0xCC, 0x29, 0xAC, 0x48, 0x18, 0x69, 
        0x5D, 0x45, 0x54, 0x9B, 0x05, 0x0C, 0x25, 0x63, 0xE7, 0x57, 0x0F, 0xC9, 
        0x10, 0xE8, 0xD6, 0x71, 0x06, 0xAC, 0xEB, 0x14, 0xE6, 0xF7, 0xB1, 0xCD, 
        0xFE, 0xCE, 0x74, 0xDC, 0xE0, 0xFD, 0x96, 0x78, 0xA6, 0x9C, 0x6D, 0x12, 
        0x95, 0x1D, 0x94, 0x85, 0x1C, 0x64, 0x1F, 0x41, 0xE7, 0x76, 0xA1, 0x3B, 
        0x80, 0xF8, 0x97, 0xD7, 0x7E, 0x47, 0x22, 0x9E, 0x16, 0x7A, 0x64, 0x41, 
        0xEA, 0x19, 0xE6, 0x54, 0x9F, 0x50, 0xA3, 0x47, 0xE8, 0xFD, 0xE7, 0x12, 
        0x53, 0xC9, 0xE3, 0xA1, 0x9E, 0x90, 0xA8, 0x99, 0xF8, 0x62, 0x48, 0x61, 
        0xE3, 0x5F, 0x45, 0xBE, 0x41, 0x3A, 0x44, 0xDA, 0x1D, 0x08, 0xF3, 0x09, 
        0x6F, 0x36, 0x41, 0x26, 0xF5, 0x4A, 0xC9, 0x0B, 0xCD, 0x5F, 0x06, 0x02, 
        0x9B, 0x22, 0x62, 0x5D, 0x97, 0x94, 0x6D, 0x2E, 0x81, 0x6F, 0x84, 0xDC, 
        0x26, 0x9E, 0x89, 0x5B, 0x9A, 0x10, 0x0A, 0xC1, 0xFA, 0x93, 0xF9, 0x5F, 
        0xFD, 0x73, 0x07, 0x5B, 0xDD, 0x20, 0x92, 0x74, 0x02, 0x7C, 0x8C, 0x1E, 
        0xFD, 0xA6, 0x17, 0x71, 0x84, 0x6C, 0xA0, 0x2C, 0x15, 0x81, 0xB7, 0x64, 
        0xBF, 0xCC, 0xDC, 0x35, 0x59, 0xEB, 0x44, 0x43, 0x55, 0xE7, 0x75, 0x38, 
        0x20, 0xFF, 0x23, 0x86, 0x60, 0xE2, 0x88, 0x64, 0x8A, 0x1C, 0x97, 0xEC, 
        0x33, 0x8A, 0x4E, 0x23, 0x0A, 0x8D, 0x31, 0x4C, 0xD6, 0x0D, 0x19, 0x7A, 
        0x01, 0x01, 0x06, 0x62, 0xE9, 0xA5, 0xFE, 0x62, 0xDD, 0xD3, 0x5C, 0x36, 
        0x97, 0x8C, 0x7E, 0x69, 0x51, 0x80, 0x34, 0x18, 0xDF, 0x24, 0xD9, 0xDF, 
        0x67, 0x57, 0xDB, 0xE3, 0xC1, 0x18, 0x36, 0x85, 0x55, 0x95, 0xB1, 0x99, 
        0x04, 0xDC, 0x47, 0x68, 0xC3, 0x11, 0x21, 0xAF, 0x19, 0xC5, 0xD2, 0x34, 
        0x1B, 0xF8, 0xA9, 0xB4, 0xC1, 0xF0, 0xDE, 0x80, 0xF9, 0xD2, 0xA4, 0x34, 
        0xC2, 0xF4, 0xBE, 0x94, 0x62, 0x7F, 0xD8, 0x11, 0xCD, 0x4B, 0x7D, 0x96, 
        0x12, 0x04, 0xEE, 0x0F, 0x02, 0x05, 0x6B, 0x2B, 0x00, 0x30, 0xA5, 0x30, 
        0x1A, 0x1F, 0x7C, 0xEF, 0xB3, 0xEE, 0x38, 0xF9, 0xD3, 0x42, 0xD6, 0xB0, 
        0xCC, 0x26, 0xC6, 0x5A, 0x3A, 0x87, 0x9E, 0xD0}
)