package toolkits

import "testing"

func TestIsIDCard(t *testing.T) {
	tests := []struct {
		arg     string
		want    bool
		wantErr bool
	}{
		{arg: "440253600213582", want: true, wantErr: false},
		{arg: "420921198603095113", want: true, wantErr: false},
		{arg: "110108198608221126", want: true, wantErr: false},
		{arg: "110108198608221125", want: false, wantErr: true},
		{arg: "350301198906180060", want: true, wantErr: false},
		{arg: "450881199412087728", want: true, wantErr: false},
		{arg: "411081199004235955", want: true, wantErr: false},
		{arg: "420821199206305032", want: true, wantErr: false},
		{arg: "411424198912068072", want: true, wantErr: false},
		{arg: "450326198912241844", want: true, wantErr: false},
		{arg: "432922198008015828", want: true, wantErr: false},
		{arg: "430981198204075412", want: true, wantErr: false},
		{arg: "450881199006112350", want: true, wantErr: false},
		{arg: "450324198809231637", want: true, wantErr: false},
		{arg: "142232199211182197", want: true, wantErr: false},
		{arg: "450111197806062156", want: true, wantErr: false},
		{arg: "612422199608063226", want: true, wantErr: false},
		{arg: "450423199409290420", want: true, wantErr: false},
		{arg: "450103197912150539", want: true, wantErr: false},
		{arg: "420982198410236015", want: true, wantErr: false},
		{arg: "452123198510084657", want: true, wantErr: false},
		{arg: "52242619811105565X", want: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			got, err := IsIDCard(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsIDCard() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsIDCard() = %v, want %v", got, tt.want)
			}
		})
	}
}
