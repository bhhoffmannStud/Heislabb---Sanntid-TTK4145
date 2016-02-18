package driver

/*
#cgo CFLAGS: -std=c11
#cgo LDFLAGS: ${SRCDIR}/simelev.a /usr/lib/x86_64-linux-gnu/libphobos2.a -lpthread -lcomedi -lm
#include "elev.h"
*/
import "C"




func Elev_init() {
	C.elev_init();
}

func Elev_set_motor_direction(dirn int) {
	C.elev_set_motor_direction(C.elev_motor_direction_t(dirn))
}


func Elev_set_button_lamp(button int, floor int, value int){
 	C.elev_set_button_lamp(C.elev_button_type_t(button) , C.int(floor), C.int(value))
}


func Elev_set_floor_indicator(floor int) {
	C.elev_set_floor_indicator(C.int(floor))
}

func Elev_set_door_open_lamp(value int){
	C.elev_set_door_open_lamp(C.int(value))
}

func Elev_set_stop_lamp(value int){
	C.elev_set_stop_lamp(C.int(value))
}


func Elev_get_button_signal(button int, floor int) int {
	return int(C.elev_get_button_signal(C.elev_button_type_t(button), C.int(floor)));
}


func Elev_get_floor_sensor_signal() int{
	return int(C.elev_get_floor_sensor_signal())
}

func Elev_get_stop_signal() int{
	return int(C.elev_get_stop_signal())
}

func Elev_get_obstruction_signal() int{
	return int(C.elev_get_obstruction_signal())
}

