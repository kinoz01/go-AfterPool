package main

import "fmt"

type ElectionResult struct {
	Name string
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	return &initialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter += increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{
		candidateName,
		votes,
	}
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf ("%s (%d)", result.Name, result.Votes)
	// Remember that the dot operator works the same
	// with a pointer to a struct as it does with the
	// actual struct. You don't have to explicitly
	// write (*result).
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate]--
}


func main() {
	// task 1: 
	counter := NewVoteCounter(2)
	fmt.Println(*counter)

	// task 2: 
	votes := 4
	voteCounter := &votes
	fmt.Println(VoteCount(voteCounter))
	var nilVoteCounter *int
	fmt.Println(VoteCount(nilVoteCounter))

	// task 3: 
	votes = 3
	IncrementVoteCount(voteCounter, 2)
	fmt.Println(votes, voteCounter, *voteCounter)

	// task 4: 
	fmt.Println(*NewElectionResult("Peter", 3))
	fmt.Println(NewElectionResult("Peter", 3))
	fmt.Println(NewElectionResult("Peter", 3).Name)
	fmt.Println(NewElectionResult("Peter", 3).Votes)

	// task 5:
	result := &ElectionResult{
		Name: "John",
		Votes: 32,
	}
	fmt.Println(DisplayResult(result))

	// task 6:
	var finalResults = map[string]int{
		"Mary":  10,
		"John":  51,
	}
	
	DecrementVotesOfCandidate(finalResults, "Mary")
	
	fmt.Println(finalResults["Mary"])
}

// Election Day
