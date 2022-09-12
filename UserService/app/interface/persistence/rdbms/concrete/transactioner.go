package concrete

import "log"

func (dbconn *DBConn) Commit() error {
	return nil
}

func (dbconn *DBConn) Rollback() error {
	return nil
}

//TxEnd DB doesnt rollback, do nothing here
func (dbconn *DBConn) TxEnd(func() error) error {
	return nil
}

func (txconn *TxConn) Commit() error {
	return txconn.DB.Commit()
}

func (txconn *TxConn) Rollback() error {
	return txconn.DB.Rollback()
}

func (txconn *TxConn) TxEnd(txFunc func() error) error {
	var err error

	// using defer to ensure transactions closes properly
	defer func() {
		if p := recover(); p != nil {
			log.Println("found p and rollback:", p)
			txconn.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			log.Println("found err and rollback: ", err)
			txconn.Rollback() // err is non-nil; don't change
		} else {
			log.Println("commit")
			err = txconn.Commit() // if Commit return error, update err with commit err
		}
	}()

	err = txFunc()
	return err
}
