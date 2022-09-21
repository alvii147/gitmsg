SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

PYTHON_VENV_NAME=.venv
PYTHON_VENV_PATH=$SCRIPT_DIR/gui/$PYTHON_VENV_NAME

rm -rf $PYTHON_VENV

if [ ! -z `which python3` ]; then
    PYTHON_CMD_PATH=`which python3`
elif [ ! -z `which python` ]; then
    PYTHON_CMD_PATH=`which python`
else
    echo "Error: python command not found, please install Python first" >&2
    exit 1
fi

cd $SCRIPT_DIR/gui/ && $PYTHON_CMD_PATH -m venv .venv && cd -

if [ -d $PYTHON_VENV_PATH/bin ]; then
    PYTHON_VENV_EXEC_PATH=$PYTHON_VENV_PATH/bin
elif [ -d $PYTHON_VENV_PATH/Scripts ]; then
    PYTHON_VENV_EXEC_PATH=$PYTHON_VENV_PATH/Scripts
else
    echo "Error: Python venv binaries not found" >&2
    exit 1
fi

if [ -f $PYTHON_VENV_EXEC_PATH/python3 ]; then
    PYTHON_CMD_PATH=$PYTHON_VENV_EXEC_PATH/python3
elif [ -f $PYTHON_VENV_EXEC_PATH/python ]; then
    PYTHON_CMD_PATH=$PYTHON_VENV_EXEC_PATH/python
else
    echo "Error: python command not found" >&2
    exit 1
fi

source $PYTHON_VENV_EXEC_PATH/activate

if [ ! -z `which pip3` ]; then
    PIP_CMD_PATH=`which pip3`
elif [ ! -z `which pip` ]; then
    PIP_CMD_PATH=`which pip`
else
    echo "Error: pip command not found, please install Python package manager first" >&2
    exit 1
fi

if [ -z $VIRTUAL_ENV ]; then
    echo "Error: Python virtual environment setup failed" >&2
    exit 1
fi

$PIP_CMD_PATH install -r $SCRIPT_DIR/gui/requirements.txt
$PIP_CMD_PATH freeze

GO_EXEC_PATH=`which go`
if [ -z $GO_EXEC_PATH ]; then
    echo "Go command not found, please install Go first" >&2
    exit 1
fi

cd $SCRIPT_DIR/cli/ && go build && cd -
