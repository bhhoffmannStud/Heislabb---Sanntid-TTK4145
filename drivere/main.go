package main
/*
#cgo LDFLAGS: -lcomedi -lm
#include "elev.h"
*/
import "C"

func motor_direction() {
	C.elev_set_motor_direction(1)

}

func init() {
	C.elev_init()

}


func main(){
	init();
	motor_direction()
	
	//return int(C.random())
}