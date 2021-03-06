package raft

import (
	"fmt"
	//"strconv"
	"math/rand"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Started Testing. This should take about 3 minutes.\n")
	var svrArr []raftServer

	for i := 1; i < 6; i++ {
		svrArr = append(svrArr, NewServer(i, "config.txt"))
	}

	//test 1 - To check if leader is elected at start. Start and close all servers in a loop for 5 times to check if election is successful at start. Leader must be elected in 2 second time

	for j := 0; j < 5; j++ {
		time.Sleep(3 * time.Second)
		check := false
		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				check = true
				break
			}
		}
		if check == false {

			t.Error("Error: Cant elect leader at start")
		}

		for i := 0; i < 5; i++ {
			svrArr[i].Close()
		}

		for i := 0; i < 5; i++ {
			svrArr[i].Start()

		}
	}
	fmt.Printf("Test 1 Successful\n")

	//test 2 - Close the leader every time and check if a new leader is elected. Then revive the leader. Repeat check 5 times.

	for j := 0; j < 5; j++ {
		time.Sleep(3 * time.Second)
		check := false
		leader := -1

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				svrArr[i].Close()
				break
			}
		}

		time.Sleep(3 * time.Second)

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				if i != leader {

					check = true
					break
				}
			}
		}

		if check == false {
			t.Error("Error: Cant elect leader after 1 server is down")
		}

		svrArr[leader].Start()
	}
	fmt.Printf("Test 2 Successful\n")

	//test 3 - close any one node at random
	for j := 0; j < 5; j++ {
		time.Sleep(2 * time.Second)
		check := false
		random := rand.Intn(5)
		svrArr[random].Close()

		time.Sleep(3 * time.Second)

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				if i != random {
					check = true
					break
				}
			}
		}

		if check == false {
			t.Error("Error: Cant elect leader after 1 server is down")
		}

		svrArr[random].Start()
	}
	fmt.Printf("Test 3 Successful\n")

	//test 4 - close any two nodes at random
	for j := 0; j < 5; j++ {
		time.Sleep(3 * time.Second)
		check := false
		random := rand.Intn(5)
		random2 := rand.Intn(5)
		if random == random2 {
			random2 = (random2 + 1) % 5
		}
		svrArr[random].Close()
		svrArr[random2].Close()

		time.Sleep(3 * time.Second)

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				if (i != random) && (i != random2) {
					check = true
					break
				}
			}
		}

		if check == false {
			t.Error("Error: Cant elect leader after 2 servers are down")
		}

		svrArr[random].Start()
		svrArr[random2].Start()
	}
	fmt.Printf("Test 4 Successful\n")

	//test 5 - close any three nodes, one of them is leader. No leader should be present
	for j := 0; j < 5; j++ {
		time.Sleep(3 * time.Second)
		check := true
		leader := -1

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				break
			}
		}

		random := rand.Intn(5)
		random2 := rand.Intn(5)
		random3 := leader
		if random == random2 {
			random2 = (random2 + 1) % 5
		}
		if random2 < random {
			temp := random
			random = random2
			random2 = temp
		}
		if random3 == random {
			random3 = (random3 + 1) % 5
		}
		if random3 == random2 {
			random3 = (random3 + 1) % 5
		}
		if random != leader {
			svrArr[random].Close()
		}
		if random2 != leader {
			svrArr[random2].Close()
		}
		if random3 != leader {
			svrArr[random3].Close()
		}
		svrArr[leader].Close()
		//fmt.Printf("%d - random\n",random)
		//fmt.Printf("%d - random2\n",random2)
		//fmt.Printf("%d - random3\n",random3)
		//fmt.Printf("%d - leader\n",leader)
		time.Sleep(3 * time.Second)

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				if (i != random) && (i != random2) && (i != random3) {
					check = false
					//fmt.Printf("Server %d is leader!\n",i)
					break
				}
			}
		}

		if check == false {
			t.Error("Error: New leader present after 3 servers are down")
		}

		svrArr[random].Start()
		svrArr[random2].Start()
		svrArr[random3].Start()
	}
	fmt.Printf("Test 5 Successful\n")
	fmt.Printf("Successfully Tested all cases\n")
}
