PREFIX=/usr/local
TARGET=swp

$TARGET:
    go build -x -v

clean:V:
    go clean -x

nuke:V:
    go clean -i -x
    if [ "${PREFIX}" != "${HOME}" ]; then
        echo "Removing the man page requires superuser privileges."
        SUDO=sudo
    else
        SUDO=
    fi
    $SUDO rm -f "$(manpath | sed -e 's/:..*//')"/${TARGET}.1
    if [ ! -z "$SUDO" ]; then
        $SUDO -k
    fi

install:V:swp
    go install -x -v
    if [ "${PREFIX}" != "${HOME}" ]; then
        echo "Installing the man page requires superuser privileges."
        SUDO=sudo
    else
        SUDO=
    fi
    $SUDO install -D ${TARGET}.1 "$(manpath | sed -e 's/:..*//')"/man1/
    if [ ! -z "$SUDO ]; then
        $SUDO -k
    fi

man:V:
    echo "manph: ${MANPATH}"

lint:V:
    go vet
