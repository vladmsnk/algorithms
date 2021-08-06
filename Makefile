SRCS = srcs/lists.cpp srcs/main.cpp tasks/reverse_list.cpp

HEADERS = includes/listf.h

CFLAGS = -Wall -Werror -Wextra 

CC = g++

all : output

output :
		${CC} ${CFLAGS} ${HEADERS} ${SRCS} -o output

clean :
		rm -rf output *.gch