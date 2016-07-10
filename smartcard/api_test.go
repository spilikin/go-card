package smartcard

import (
    "fmt"
    "testing"
)

var CMD_SELECT = []byte{
    0x00, 0xA4, 0x04, 0x00, 0x08,
    0x90, 0x72, 0x5A, 0x9E, 0x3B, 0x10, 0x70, 0xAA,
}

var CMD_10 = []byte{
    0x00, 0x10, 0x00, 0x00, 0x0B,
}

func printHex(buffer []byte) {
    for _, b := range buffer {
        fmt.Printf("%02x", b)
    }
    fmt.Println("")
}

func TestEstablishReleaseContext(t *testing.T) {
    fmt.Println("\n==============================")
    fmt.Println("Test establish/release Context")
    fmt.Println("==============================\n")
    ctx, err := EstablishContext()
    if err != nil { t.Error(err); return }
    err = ctx.Release()
    if err != nil { t.Error(err); return }
    fmt.Println("OK")
}

func TestListReaders(t *testing.T) {
    fmt.Println("\n=================")
    fmt.Println("Test list readers")
    fmt.Println("=================\n")
    ctx, err := EstablishContext()
    if err != nil { t.Error(err); return }
    defer ctx.Release()
    readers, err := ctx.ListReaders()
    if err != nil { t.Error(err); return }
    for _, reader := range readers {
        fmt.Println(reader.Name())
        fmt.Printf("- Card present: %t\n\n", reader.IsCardPresent())
    }
}

/*
func TestHighLevelAPI(t *testing.T) {
    fmt.Println("\n===================")
    fmt.Println("High Level API Test")
    fmt.Println("===================\n")

    fmt.Println("\nEstablish Context")
    fmt.Println("-----------------\n")
    ctx, err := EstablishContext()
    if err != nil { t.Error(err); return }
    defer ctx.Release()
    fmt.Println("OK")

    fmt.Println("\nWait for card present")
    fmt.Println("---------------------\n")
    reader, err := ctx.WaitForCardPresent()
    if err != nil { t.Error(err); return }
    fmt.Println("OK")

    fmt.Println("\nConnect to card")
    fmt.Println("---------------\n")
    card, err := reader.Connect()
    if err != nil { t.Error(err); return }
    defer card.Disconnect()
    fmt.Print("ATR: ")
    printHex(card.ATR())

    fmt.Println("\nSelect applet")
    fmt.Println("-------------\n")
    printHex(CMD_SELECT)
    response, err := card.Transmit(CMD_SELECT)
    if err != nil { t.Error(err); return }
    printHex(response)

    fmt.Println("\nSend CMD 10")
    fmt.Println("-----------\n")
    printHex(CMD_10)
    response, err = card.Transmit(CMD_10)
    if err != nil { t.Error(err); return }
    printHex(response)
    fmt.Printf("Quoth the Applet, \"%s\"\n", string(response[:len(response)-2]))
}
*/