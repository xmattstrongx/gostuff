package main

import (
	"testing"
)

func TestGateway(t *testing.T) {

	s := PayData{name: "Matt", date: 1460125296, payrate: 1000, paytime: 80}

	Gateway(s)
	t.Fail()
}

func TestAdapter(t *testing.T) {

	s := PayData{name: "Matt", date: 1460125296, payrate: 1000, paytime: 80}

	Adapter(s)
	t.Fail()
}

func TestAggregate(t *testing.T) {

	s := PayData{name: "Matt", date: 1460125296, payrate: 1000, paytime: 80}

	Aggregate(s)
	t.Fail()
}

func TestView(t *testing.T) {

	s := PayData{name: "Matt", date: 1460125296, payrate: 1000, paytime: 80}

	View(s)
	t.Fail()
}
