package jwt

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_jwt(t *testing.T) {
	type args struct {
		uid        int64
		username   string
		expireTime int64
		sleep      int64
	}
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr error
	}{
		{
			name: "generate token",
			args: args{
				uid:      1,
				username: "wx",
			},
			want: &Claims{
				Uid:      1,
				UserName: "wx",
			},
			wantErr: nil,
		},
		{
			name: "token parse fail expire",
			args: args{
				uid:        1,
				username:   "wx",
				expireTime: time.Now().Add(time.Second * 2).Unix(),
				sleep:      3,
			},
			want: &Claims{
				Uid:      1,
				UserName: "wx",
			},
			wantErr: ErrTokenExpired,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jwt{
				key:        []byte("33446a9dcf9ea060a0a6532b166da32f304af0de"),
				maxRefresh: 10,
			}
			if tt.args.expireTime > 0 {
				j.expireTime = tt.args.expireTime
			}
			tokenString := j.GenerateToken(tt.args.uid, tt.args.username)
			claims, err := j.ParseTokenBy(tokenString)
			assert.Nil(t, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.want.Uid, claims.Uid)
			assert.Equal(t, tt.want.UserName, claims.UserName)

			if j.expireTime > 0 {
				time.Sleep(time.Duration(tt.args.sleep) * time.Second)
			}
			_, err = j.ParseTokenBy(tokenString)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
