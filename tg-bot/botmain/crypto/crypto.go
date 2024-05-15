package crypto

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"

	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent"
	"github.com/Simplewallethq/simple-wallet-tg-bot/tg-bot/ent/user"
	"github.com/make-software/casper-go-sdk/types/keypair"
	csprsecp "github.com/make-software/casper-go-sdk/types/keypair/secp256k1"
	"github.com/pkg/errors"
)

type Crypto struct {
	DB            *ent.Client
	TempPassStore map[int64]string
	Chain         string
	PK_SALT       string
}

func NewCrypto(DB *ent.Client, salt string, chain string) *Crypto {
	return &Crypto{
		DB:            DB,
		TempPassStore: make(map[int64]string),
		Chain:         chain,
		PK_SALT:       salt,
	}
}

func (c *Crypto) GenerateNewUserWallet(id int64, password string) ([]byte, error) {
	pair, pemData, err := c.GenerateEd25519Pair()
	if err != nil {
		return nil, err
	}
	u, err := c.DB.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}
	err = u.Update().SetPublicKey(pair.PublicKey().String()).Exec(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed update user")
	}
	key := deriveKey(password, []byte(c.PK_SALT))
	encryptMsg, err := encrypt(key, pemData)
	if err != nil {
		return nil, errors.Wrap(err, "failed encrypt pemData")
	}
	err = c.DB.PrivateKeys.Create().SetOwner(u).SetPrivateKey(encryptMsg).Exec(context.Background())
	//err = u.Update().SetPrivatKey(encryptedPem).Exec(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed update user")
	}
	// DEC_PEM, err := decrypt(key, encryptMsg)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed decrypt pemData")
	// }
	// log.Println("DecPEM: ", string(DEC_PEM))
	return pemData, nil
}

func (c *Crypto) NewWalletFromPEM(id int64, content []byte) error {
	block, _ := pem.Decode(content)
	if block == nil {
		return errors.New("failed decode pem")
	}
	var err error
	ed := []byte{48, 46, 2}
	secp := []byte{48, 116, 2}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("failed decode pem")
		}
	}()
	var pair *keypair.PrivateKey
	if bytes.Equal(block.Bytes[:3], ed) {
		log.Println("ed25519")
		pair, err = c.Ed25519PairFromPEM(content)
		if err != nil {
			return err
		}
	} else if bytes.Equal(block.Bytes[:3], secp) {
		log.Println("secp256k1")
		pair, err = c.SecpPairFromPEM(content)
		if err != nil {
			return err
		}
	} else {
		return errors.New("unknown key type")
	}

	u, err := c.DB.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed get user")
	}
	err = u.Update().SetPublicKey(pair.PublicKey().String()).Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}
	var userPass string
	if p, ok := c.TempPassStore[id]; ok {
		userPass = p
	} else {
		return errors.New("password not found")
	}

	key := deriveKey(userPass, []byte(c.PK_SALT))
	encryptMsg, err := encrypt(key, content)
	if err != nil {
		return errors.Wrap(err, "failed encrypt pem")
	}
	err = c.DB.PrivateKeys.Create().SetOwner(u).SetPrivateKey(encryptMsg).Exec(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed update user")
	}

	//delete pass from store
	delete(c.TempPassStore, id)
	return nil
}

func (c *Crypto) NewKeypairFromPEM(id int64, content []byte) (*keypair.PrivateKey, error) {
	block, _ := pem.Decode(content)
	if block == nil {
		return nil, errors.New("failed decode pem")
	}
	var err error
	ed := []byte{48, 46, 2}
	secp := []byte{48, 116, 2}
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("failed decode pem")
		}
	}()
	var pair *keypair.PrivateKey
	if bytes.Equal(block.Bytes[:3], ed) {
		log.Println("ed25519")
		pair, err = c.Ed25519PairFromPEM(content)
		if err != nil {
			return nil, err
		}
	} else if bytes.Equal(block.Bytes[:3], secp) {
		log.Println("secp256k1")
		pair, err = c.SecpPairFromPEM(content)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("unknown key type")
	}

	return pair, nil
}

func (c *Crypto) GenerateNewUserWalletWithoutStorePK(id int64, password string) ([]byte, error) {
	pair, pemData, err := c.GenerateEd25519Pair()
	if err != nil {
		return nil, err
	}
	u, err := c.DB.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}
	err = u.Update().SetPublicKey(pair.PublicKey().String()).Exec(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed update user")
	}
	// dk := pbkdf2.Key([]byte(password), []byte(PK_SALT), 4096, 32, sha1.New)
	// aes, err := aes.NewCipher(dk[:])
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed create cipher")
	// }
	// encryptedPem := make([]byte, len(pemData))
	// aes.Encrypt(encryptedPem, []byte(pemData))
	// err = u.Update().SetPrivatKey(encryptedPem).Exec(context.Background())
	// if err != nil {
	// 	return nil, errors.Wrap(err, "failed update user")
	// }
	return pemData, nil
}

func (c *Crypto) SaveUserPassword(id int64, password string) error {
	c.TempPassStore[id] = password
	return nil
}

func TestParswPem() {

	fileData, err := os.ReadFile("secp.pem")
	if err != nil {
		panic(err)
	}
	one, two := pem.Decode(fileData)
	log.Println(len(one.Bytes))
	log.Println(two)

}

func GenerateEd25519Pair() (keypair.PrivateKey, []byte, error) {
	var (
		err   error
		b     []byte
		block *pem.Block
		priv  ed25519.PrivateKey
	)

	_, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Generation error : %s", err)
		os.Exit(1)
	}

	b, err = x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return keypair.PrivateKey{}, nil, err
	}

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}
	pemContent := pem.EncodeToMemory(block)
	//pair, err := keypair.NewPrivateKeyED25518FromPEM(pemContent)
	pair, err := keypair.NewPrivateKeyFromPEM(pemContent, keypair.ED25519)
	if err != nil {
		return keypair.PrivateKey{}, nil, err
	}
	log.Println("Public key: ", pair.PublicKey())
	return pair, pemContent, nil
}

func GenerateSecp256k1Pair() (keypair.PrivateKey, []byte, error) {
	priv, _, err := csprsecp.NewPemPair()
	if err != nil {
		return keypair.PrivateKey{}, nil, err
	}
	pair, err := keypair.NewPrivateKeyFromPEM(priv, keypair.SECP256K1)
	if err != nil {
		return keypair.PrivateKey{}, nil, err
	}
	//log.Println("PEM Public key: ", string(pub))
	//log.Println("PEM Private key: ", string(priv))
	log.Println("Public key: ", pair.PublicKey())
	return pair, priv, err
}

// GenerateSaveEd25519  example how to generate and save ed25519 key pair to pem files
func GenerateSaveEd25519(fb string) error {

	var (
		err   error
		b     []byte
		block *pem.Block
		pub   ed25519.PublicKey
		priv  ed25519.PrivateKey
	)

	pub, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Generation error : %s", err)
		os.Exit(1)
	}

	b, err = x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}

	err = os.WriteFile(fb, pem.EncodeToMemory(block), 0600)
	if err != nil {
		return err
	}

	// public key
	b, err = x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}

	fileName := fb + ".pub"
	err = os.WriteFile(fileName, pem.EncodeToMemory(block), 0644)
	return err

}

func (c *Crypto) GenerateEd25519Pair() (*keypair.PrivateKey, []byte, error) {
	//good
	var (
		err   error
		b     []byte
		block *pem.Block
		priv  ed25519.PrivateKey
	)

	_, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Generation error : %s", err)
		os.Exit(1)
	}

	b, err = x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return nil, nil, err
	}

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}
	pemContent := pem.EncodeToMemory(block)
	pair, err := keypair.NewPrivateKeyFromPEM(pemContent, keypair.ED25519)
	if err != nil {
		return nil, nil, err
	}
	return &pair, pemContent, nil
}

func (c *Crypto) Ed25519PairFromPEM(content []byte) (*keypair.PrivateKey, error) {

	pair, err := keypair.NewPrivateKeyFromPEM(content, keypair.ED25519)
	if err != nil {
		return nil, err
	}
	return &pair, nil
}

func (c *Crypto) SecpPairFromPEM(content []byte) (*keypair.PrivateKey, error) {
	pair, err := keypair.NewPrivateKeyFromPEM(content, keypair.SECP256K1)
	if err != nil {
		return nil, err
	}
	return &pair, nil
}

func (c *Crypto) DecodePemFromDB(uid int64, password string) (*keypair.PrivateKey, error) {
	u, err := c.DB.User.Query().Where(user.ID(uid)).Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}
	PKStore, err := u.QueryPrivateKey().Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}
	key := deriveKey(password, []byte(c.PK_SALT))
	decMsg, err := decrypt(key, PKStore.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed decrypt pem")
	}
	block, _ := pem.Decode(decMsg)
	if block == nil {
		return nil, errors.New("failed decode pem")
	}
	ed := []byte{48, 46, 2}
	secp := []byte{48, 116, 2}
	var pair *keypair.PrivateKey
	if bytes.Equal(block.Bytes[:3], ed) {
		log.Println("ed25519")
		pair, err = c.Ed25519PairFromPEM(decMsg)
		if err != nil {
			return nil, err
		}
	} else if bytes.Equal(block.Bytes[:3], secp) {
		log.Println("secp256k1")
		pair, err = c.SecpPairFromPEM(decMsg)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("unknown key type")
	}
	log.Println("PubKey ", pair.PublicKey())
	return pair, nil
}

func (c *Crypto) PemFromDB(uid int64, password string) ([]byte, error) {
	u, err := c.DB.User.Query().Where(user.ID(uid)).Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}
	PKStore, err := u.QueryPrivateKey().Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed get user")
	}
	key := deriveKey(password, []byte(c.PK_SALT))
	decMsg, err := decrypt(key, PKStore.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed decrypt pem")
	}
	block, _ := pem.Decode(decMsg)
	if block == nil {
		return nil, errors.New("failed decode pem")
	}
	ed := []byte{48, 46, 2}
	secp := []byte{48, 116, 2}
	var pk []byte
	if bytes.Equal(block.Bytes[:3], ed) {
		log.Println("ed25519")
		pk = decMsg
	} else if bytes.Equal(block.Bytes[:3], secp) {
		log.Println("secp256k1")
		pk = decMsg
	} else {
		return nil, errors.New("unknown key type")
	}
	return pk, nil
}
