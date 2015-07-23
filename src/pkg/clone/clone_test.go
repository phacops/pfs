package clone

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasic(t *testing.T) {
	dirPath, err := ioutil.TempDir("", "pachyderm")
	require.NoError(t, err)
	err = GithubClone(
		dirPath,
		"pachyderm",
		"pachyderm",
		"master",
		"11eb4c6e1945beb1e6ce3e878ed2cb6c24ee8bf1",
		"",
	)
	fmt.Println(dirPath)
	require.NoError(t, err)
	//data, err := ioutil.ReadFile(filepath.Join(dirPath, "README.md"))
	//require.NoError(t, err)
	//require.Equal(
	//t,
	//`pfs
	//===

	//The Pachyderm Filesystem`,
	//string(data),
	//)
}