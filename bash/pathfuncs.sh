source "${BASH_LIBS}/strfuncs.sh"

###
# The following functions operate on the PATH variable, or, if specified, a
# different environment variable with similar semantics.

function lspath() {
# Displays the members of the path variable on separate lines
	local pathlistname=$1
	[[ -z $pathlistname ]] && pathlistname="PATH"

	local pathlist=${!pathlistname}

	echo $pathlist | tr ':' '\n' | awk 'NF > 0' ||
        echo "ERROR in lspath(): pathlistname=$pathlistname pathlist=$pathlist" 1>&2
}

function _removePath() {
# if provided, the argument specifies the name of the variable to
# operate on. Otherwise, it defaults to $PATH.
# DOES NOT export! echoes the result to stdout.
    local pathlistname=$2
    [[ -z $pathlistname ]] && pathlistname="PATH"

    local pathlist="${!pathlistname}"

    local result=''
    for p in `lspath "${pathlistname}" | tr ' ' '\a'`; do
        p="`echo "$p" | tr '\a' ' '`"
        [[ $p != $1 ]] && if [[ -n ${result} ]]; then
            result="${result}:$p"
        else
            result="$p"
        fi
    done
    echo "$result"
}

function removePath() {
# the first argument is the string to remove from the path list.

# if provided, the second argument specifies the name of the variable to
# operate on. Otherwise, it defaults to $PATH.
    local pathlistname=$2
    [[ -z $pathlistname ]] && pathlistname="PATH"

    local pathlist=${!pathlistname}
    pathlist="`_removePath "$1" "${pathlistname}"`"

    export $pathlistname="${pathlist}"
}

function prependPath() {
# the first argument is the string to prepend to the path list.

# if provided, the second argument specifies the name of the variable to
# operate on. Otherwise, it defaults to $PATH.
    [[ -d "$1" ]] || return 1
    local pathlistname=$2
    [[ -z $pathlistname ]] && pathlistname="PATH"

    local pathlist=${!pathlistname}
    pathlist="`_removePath "${1}" "${pathlistname}"`"

    if [[ -z $pathlist ]]; then
        export $pathlistname="${1}"
    else
        export $pathlistname="${1}:${pathlist}"
    fi
}

function appendPath() {
# the first argument is the string to append to the path list.

# if provided, the second argument specifies the name of the variable to
# operate on. Otherwise, it defaults to $PATH.
    [[ -d "$1" ]] || return 1
    local pathlistname=$2
    [[ -z $pathlistname ]] && pathlistname="PATH"

    local pathlist=${!pathlistname}
    pathlist="`_removePath "${1}" "${pathlistname}"`"

    if [[ -z $pathlist ]]; then
        export $pathlistname="${1}"
    else
        export $pathlistname="${pathlist}:${1}"
    fi
}


###
# The following functions operate on the provided path string. They do not
# access the PATH environment variable.

function normpath() {
# removes extra slashes and resolves as many dots as possible in the provided
# path.
# Ported from Python.
    local path="$1"

    [ -z "$path" ] && echo '.' && return 0

    # POSIX allows one or two initial slashes, but treats three or more
    # as a single slash
    local initial_slashes=0
    startswith / "$path" && initial_slashes=1
    startswith // "$path" && ! startswith /// "$path" && initial_slashes=2

    local result=""
    for comp in `splitstr "$path" /`; do
        [ -z "$comp" ] || [ "$comp" = "." ] && continue
        if [ "$comp" != ".." ] ||
            ([ $initial_slashes -eq 0 ] && [ -z "$result" ]) ||
            [ "$result" = ".." ] ||
            [ "`echo "$result" | xargs | tail -c4`" = "/.." ]
        then
            [ -n "$result" ] && result="$result/"
            result="$result$comp"
        elif [ -n "$result" ]; then
            echo "$result" | grep / >/dev/null && result="${result%/*}" || result=""
        fi
    done

    result="`repeatstr / $initial_slashes`$result"
    [ -z "$result" ] && result="."
    echo "$result"
}

function isabspath() {
# Ported from Python.
    [ "`echo ${1} | head -c1`" = "/" ]
}

function expanduser() {
# expands the tilde in the provided path.
# Ported from Python.
    local path="$1"

    if ! startswith '~' "$path"; then
        echo "$path"
        return 0
    fi

    local userpart="${path%%\/*}"
    if [ "$userpart" = "~" ]; then
        local username="`whoami`"
    else
        local username="${userpart#*~}"
    fi
    rest="${path#"$userpart"}"

    # This eval is safe, because printf %q escapes $userpart.
    eval local userhome="$(printf %q "$userpart")"
    if [ -z "$userhome" ]; then
        userhome="$userpart"
    else
        userhome="${userhome%/}"
    fi

    echo "${userhome}${rest}"
}

function joinpath() {
# joins the given arguments into a path string.
# Ported from Python.
    local rv="$1"
    shift
    for sub in "$@"; do
        if [ "`echo "$sub" | head -c1`" = "/" ]; then
            rv="$sub"
        elif [ -z "$rv" ] || [ "`echo "$rv" | xargs | tail -c2`" = "/" ]; then
            rv="$rv$sub"
        else
            rv="$rv/$sub"
        fi
    done
    echo "$rv"
}

function abspath() {
# uses string operations together with $PWD to convert the input into an
# absolute path. The resulting path is not required to exist.
# Ported from Python.
    local path="$1"

    if ! isabspath "$path"; then
        path="`joinpath "$PWD" "$path"`"
    fi

    normpath "$path"
}

function relpath() {
# resolves and outputs the given path against $PWD (or, if provided, a second
# path).
    normpath "$(joinpath "$(expanduser "${2:-$PWD}")" "$1")"
}

