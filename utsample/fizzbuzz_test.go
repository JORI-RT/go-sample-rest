package utsample

import "testing"

func Test_fizzbuzz(t *testing.T) {
	type args struct {
		i int16
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "3を渡すとfizzを返す",
			args: args{
				i: 3,
			},
			want: "fizz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzbuzz(tt.args.i); got != tt.want {
				t.Errorf("fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
