package employee

import "fmt"

type Employee struct{
	FistName string
	LastName string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining(){
	fmt.Printf("%s %s has %d leaves remaining",e.FistName,e.LastName,(e.TotalLeaves-e.LeavesTaken))
}
