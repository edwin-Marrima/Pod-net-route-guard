package handler

import (
	"github.com/edwin-Marrima/Pod-net-route-guard/internal/schema"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_readConfiguration(t *testing.T) {
	type args struct {
		configFilePath string
	}
	tests := []struct {
		name    string
		args    args
		want    *schema.Config
		wantErr bool
	}{
		{
			name: "Return error when provided configuration file does not exist",
			args: args{
				configFilePath: "./fixtures/does-not-exist.yaml",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Must return schema.Rule when provided file contain a valid schema",
			args: args{
				configFilePath: "./test_data/valid-schema.yaml",
			},
			want: &schema.Config{
				Rules: schema.Rules{
					NAT: []schema.NAT{
						{
							Name: "redirect-rule-1",
							Source: schema.Source{
								IP:       "192.168.1.10",
								Port:     "8080",
								Protocol: "tcp",
							},
							Destination: schema.Destination{
								IP:   "10.0.0.5",
								Port: "443",
							},
							Action: schema.Action{
								RedirectTo: schema.RedirectTo{
									Port: "15002",
								},
							},
						},
						{
							Name: "redirect-rule-2",
							Source: schema.Source{
								IP:       "10.0.0.5",
								Port:     "443",
								Protocol: "tcp",
							},
							Action: schema.Action{
								RedirectTo: schema.RedirectTo{
									Port: "15002",
								},
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readConfiguration(tt.args.configFilePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}