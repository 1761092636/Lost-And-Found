package eth

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func NewAccount() {
	ks := keystore.NewKeyStore("/path/to/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)
	fmt.Println(am)
	// Create a new account with the specified encryption passphrase.
	newAcc, _ := ks.NewAccount("123456")
	fmt.Println(newAcc)

	// Export the newly created account with a different passphrase. The returned
	// data from this method invocation is a JSON encoded, encrypted key-file.
	jsonAcc, _ := ks.Export(newAcc, "Creation password", "Export password")

	// Update the passphrase on the account created above inside the local keystore.
	_ = ks.Update(newAcc, "Creation password", "Update password")

	// Delete the account updated above from the local keystore.
	_ = ks.Delete(newAcc, "Update password")

	// Import back the account we've exported (and then deleted) above with yet
	// again a fresh passphrase.
	impAcc, _ := ks.Import(jsonAcc, "Export password", "Import password")
	fmt.Println(impAcc)
}
