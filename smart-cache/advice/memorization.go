package advice

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/wesovilabs/beyond/api"
	"github.com/wesovilabs/beyond/api/context"
	"strings"
	"time"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

type Memorization struct {
	expirationTimeInSeconds time.Duration
}

func (m *Memorization) Before(ctx *context.BeyondContext) {
	hash := hash(ctx.Pkg(),ctx.Function(),ctx.Params())
	val, found := c.Get(hash)
	if !found {
		ctx.Set("hash", hash)
		return
	}
	ctx.Exit()
	if val != nil {
		results := val.(*context.Args)
		ctx.SetResults(results)
	}
}

func (m *Memorization) Returning(ctx *context.BeyondContext) {
	hash := ctx.Get("hash").(string)
	c.Set(hash, ctx.Results(), m.expirationTimeInSeconds)
}

func Memorize(expirationTime int) func() api.Around {
	return func() api.Around {
		return &Memorization{
			time.Duration(expirationTime) * time.Second,
		}
	}
}


func hashValue(value interface{}) string {
	hasher := md5.New()
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		hasher.Write([]byte(fmt.Sprintf("%#v",value)))
	}else{
		hasher.Write(buf.Bytes())
	}
	return hex.EncodeToString(hasher.Sum(nil))
}

func hash(pkg string, function string, args *context.Args) string {

	return fmt.Sprintf("%s.%s.%s",
		hashValue(pkg),
		hashValue(function), func() string {
			argsHashes := make([]string, args.Count())
			args.ForEach(func(index int, arg *context.Arg) {
				argsHashes[index] = hashValue(arg)
			})
			return strings.Join(argsHashes, ".")
		}())
}
