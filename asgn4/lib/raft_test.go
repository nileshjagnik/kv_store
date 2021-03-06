package raft
import (
	"fmt"
	"strings"
	"strconv"
	"math/rand"
	"testing"
	"time"
)

func TestMain(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Started Testing. This should take about 2 minutes.\n")
	var svrArr []raftServer

	for i := 1; i < 6; i++ {
		svrArr = append(svrArr, NewServer(i, "config.txt"))
	}
	time.Sleep(1 * time.Second)
	/*
	//test 1 - To check if leader is elected at start. Start and close all servers in a loop for 5 times to check if election is successful at start. Leader must be elected in 2 second time

	for j := 0; j < 1; j++ {
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
		        //fmt.Printf("Closing Server %d\n",random+1)
			svrArr[random].Close()
		}
		if random2 != leader {
		        //fmt.Printf("Closing Server %d\n",random2+1)
			svrArr[random2].Close()
		}
		if random3 != leader {
		        //fmt.Printf("Closing Server %d\n",random3+1)
			svrArr[random3].Close()
		}
		//fmt.Printf("Closing Server %d\n",leader+1)
		svrArr[leader].Close()
		//fmt.Printf("%d - random\n",random+1)
		//fmt.Printf("%d - random2\n",random2+1)
		//fmt.Printf("%d - random3\n",random3+1)
		//fmt.Printf("%d - leader\n",leader+1)
		time.Sleep(3 * time.Second)

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				if (i != random) && (i != random2) && (i != random3) {
					check = false
					fmt.Printf("Server %d is leader!\n",i+1)
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
	*/
	//test 6 - when a command is given to a non leader shows error 
	for j := 0; j < 5; j++ {
	        random := rand.Intn(5)
	        if (svrArr[random].isLeader()==true) {
                        random = (random + 1) % 5
                }
	        svrArr[random].Outbox() <- "set abc xyz"
	        time.Sleep(375 * time.Millisecond)
	        select {
		case msg := <-svrArr[random].ClientSideOutbox():
		//fmt.Printf(msg.(string)+"\n")
		        if (!strings.HasPrefix(msg.(string), "Request Failed: Leader is ")) {
		                t.Error("Error: Non Leader not showing error")
		        }
		}
	        
	}
	fmt.Printf("Test 6 Successful\n")
	
	// test 7 - Integrity test - to check if a result is produced when all servers are up
	for j := 0; j < 5; j++ {
	        leader := -1

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				break
			}
		}
		key := "abc" + strconv.Itoa(j)
		val := "xyz" + strconv.Itoa(j)
	        svrArr[leader].Outbox() <- "set " + key +" " + val
	        
	        for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				break
			}
		}
	        select {
		case <-svrArr[leader].ClientSideOutbox():
		        break
		        //fmt.Printf(msg.(string)+"\n")
		case <-time.After(10 * time.Second):
		        t.Error("Error: No Reply from Leader after 10 seconds")
		}
		time.Sleep(time.Second)
		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				break
			}
		}
		svrArr[leader].Outbox() <- "get " + key
	        select {
		case msg := <-svrArr[leader].ClientSideOutbox():
		        if msg.(string) == val {
		                break
		        } else {
		                t.Error("Error: Wrong value for key recieved")
		        }
		case <-time.After(10 * time.Second):
		        t.Error("Error: No Reply from Leader after 10 seconds")
		}
		
	        
	}
	fmt.Printf("Test 7 Successful\n")
	time.Sleep(2 * time.Second)
	/*
	// test 8 - integrity test after 1 server is down
	for j := 0; j < 5; j++ {
	        fmt.Println(j)
	        random := rand.Intn(5)
	        svrArr[random].Close()
	        time.Sleep(2 * time.Second)
	        leader := -1

		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				break
			}
		}
		key := "abc" + strconv.Itoa(j)
		val := "xyz" + strconv.Itoa(j)
	        svrArr[leader].Outbox() <- "set " + key +" " + val
	        
	        for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				break
			}
		}
	        select {
		case msg:= <-svrArr[leader].ClientSideOutbox():
		        fmt.Printf(msg.(string)+"\n")
		        break
		case <-time.After(10 * time.Second):
		        t.Error("Error: No Reply from Leader after 10 seconds")
		}
		time.Sleep(time.Second)
		for i := 0; i < 5; i++ {
			if svrArr[i].isLeader() == true {
				leader = i
				break
			}
		}
		svrArr[leader].Outbox() <- "get " + key
	        select {
		case msg := <-svrArr[leader].ClientSideOutbox():
		        if msg.(string) == val {
		                break
		        } else {
		                t.Error("Error: Wrong value for key recieved")
		        }
		case <-time.After(10 * time.Second):
		        t.Error("Error: No Reply from Leader after 10 seconds")
		}
		svrArr[random].Start()
	}
	fmt.Printf("Test 8 Successful\n")
	*/
	fmt.Printf("Successfully Tested all cases\n")
	
}
