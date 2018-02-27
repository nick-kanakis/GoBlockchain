package api

import (
	"flag"
	"os"
	"fmt"
	"log"
	"personal/GoBlockchain/blockchain"
	"personal/GoBlockchain/persistance"
)

type CLI struct{}

func (cli *CLI) validateArguments(){
	if len(os.Args) < 2{
		log.Println("Less than 2 arguments")
		cli.showHelp()
		os.Exit(1)
	}
}

func (cli *CLI) showHelp(){
	fmt.Println("Available commands:")
	fmt.Println("	-insertdata -data DATA - insert data into blockchain.")
	fmt.Println("	-printchain - Print block chain from newsest block to oldest.")
	fmt.Println("	-validatechain - validate blockchain.")
}

func (cli *CLI) Run(){
	cli.validateArguments()
	insertCmd := flag.NewFlagSet("insertdata", flag.ExitOnError)
	insertData := insertCmd.String("data", "", "The data to be stored in block")
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	validateChainCmd := flag.NewFlagSet("validateChain", flag.ExitOnError)

	switch os.Args[1]{
	
		case "-insertdata":
			err := insertCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "-printchain":
			err := printChainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "-validatechain":
			err := validateChainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		default:
			log.Printf("Argument not valid argument: %v", os.Args[1])
			cli.showHelp()
			os.Exit(1)		
	}

	if insertCmd.Parsed(){
		if *insertData == ""{
			insertCmd.Usage()
		}
		cli.insert(*insertData)
	}

	if printChainCmd.Parsed(){
		cli.printChain()
	}

	if validateChainCmd.Parsed(){
		cli.validateChain()
	}
}

func (cli *CLI) insert(data string){
	pm :=persistance.NewPersistanceManager()
	defer pm.ClosePersistanceManager()
	bc, err := blockchain.NewBlockchain(pm)
	if err != nil {
		log.Panic(err)
	}
	bc.AddBlock([]byte(data))
}

func (cli *CLI) printChain(){
	pm :=persistance.NewPersistanceManager()
	defer pm.ClosePersistanceManager()
	bc, err := blockchain.NewBlockchain(pm)
	if err != nil {
		log.Panic(err)
	}
	bc.PrintChain()
}

func (cli *CLI) validateChain(){
	pm :=persistance.NewPersistanceManager()
	defer pm.ClosePersistanceManager()
	bc, err := blockchain.NewBlockchain(pm)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bc.ValidateChain())
}





