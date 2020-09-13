package utsample

import "testing"

func Test_fizzbuzz(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "3を渡すとfizzを返す",
			args: args{
				i: 3,
			},
			want: "fizz",
		},
		{
			name: "5を渡すとbuzzを返す",
			args: args{
				i: 5,
			},
			want: "buzz",
		},
		{
			name: "1を渡すと1を返す",
			args: args{
				i: 1,
			},
			want: "1",
		},
		{
			name: "15を渡すとfizzBuzzを返す",
			args: args{
				i: 15,
			},
			want: "fizzBuzz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.i); got != tt.want {
				t.Errorf("fizzbuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
