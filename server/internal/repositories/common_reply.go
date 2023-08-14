package repositories

type RowsAffected = int64

type CommonUpdateAndDeleteReply struct {
	RowsAffected RowsAffected
}
