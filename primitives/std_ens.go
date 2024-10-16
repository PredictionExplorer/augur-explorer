package primitives

import (
	"strings"

	"golang.org/x/crypto/sha3"
	"golang.org/x/net/idna"
)
var (
	p = idna.New(idna.MapForLookup(), idna.StrictDomainName(false), idna.Transitional(false))
)
// Normalize normalizes a name according to the ENS rules
func Normalize(input string) (output string, err error) {
	output, err = p.ToUnicode(input)
	if err != nil {
		return
	}
	// If the name started with a period then ToUnicode() removes it, but we want to keep it
	if strings.HasPrefix(input, ".") && !strings.HasPrefix(output, ".") {
		output = "." + output
	}
	return
}
// LabelHash generates a simple hash for a piece of a name.
func LabelHash(label string) (hash [32]byte, err error) {
	normalizedLabel, err := Normalize(label)
	if err != nil {
		return
	}

	sha := sha3.NewLegacyKeccak256()
	if _, err = sha.Write([]byte(normalizedLabel)); err != nil {
		return
	}
	sha.Sum(hash[:0])
	return
}


// NameHash generates a hash from a name that can be used to
// look up the name in ENS
func NameHash(name string) (hash [32]byte, err error) {
	if name == "" {
		return
	}
	normalizedName, err := Normalize(name)
	if err != nil {
		return
	}
	parts := strings.Split(normalizedName, ".")
	for i := len(parts) - 1; i >= 0; i-- {
		if hash, err = nameHashPart(hash, parts[i]); err != nil {
			return
		}
	}
	return
}
func nameHashPart(currentHash [32]byte, name string) (hash [32]byte, err error) {
	sha := sha3.NewLegacyKeccak256()
	if _, err = sha.Write(currentHash[:]); err != nil {
		return
	}
	nameSha := sha3.NewLegacyKeccak256()
	if _, err = nameSha.Write([]byte(name)); err != nil {
		return
	}
	nameHash := nameSha.Sum(nil)
	if _, err = sha.Write(nameHash); err != nil {
		return
	}
	sha.Sum(hash[:0])
	return
}
