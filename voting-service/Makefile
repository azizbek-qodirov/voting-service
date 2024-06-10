exp:
	export DBURL="postgres://mrbek:QodirovCoder@localhost:5432/election_voting?sslmode=disable"

mig-up:
	migrate -path db/migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path db/migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir db/migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir db/migrations -seq insert_table

	
# mig-delete:
# 	rm -r db/migrations