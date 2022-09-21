echo "Getting script directory"
SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

PYTHON_VENV_NAME=.venv
PYTHON_VENV_PATH=$SCRIPT_DIR/gui/$PYTHON_VENV_NAME

echo "Removing existing Python virtual environments"
rm -rf $PYTHON_VENV

echo "Getting python command"
PYTHON_CMD=python3
$PYTHON_CMD --version > /dev/null 2>&1
if [ ! $? -eq 0 ]; then
    PYTHON_CMD=python
fi

$PYTHON_CMD --version > /dev/null 2>&1
if [ ! $? -eq 0 ]; then
    echo "Error: python command not found, please install Python first" >&2
    exit 1
fi

echo "Found Python installation"
$PYTHON_CMD --version

echo "Creating Python virtual environment"
cd $SCRIPT_DIR/gui/ && $PYTHON_CMD -m venv .venv && cd -

echo "Locating Python virtual environment binaries"
if [ -d $PYTHON_VENV_PATH/bin ]; then
    PYTHON_VENV_EXEC_PATH=$PYTHON_VENV_PATH/bin
elif [ -d $PYTHON_VENV_PATH/Scripts ]; then
    PYTHON_VENV_EXEC_PATH=$PYTHON_VENV_PATH/Scripts
else
    echo "Error: Python venv binaries not found" >&2
    exit 1
fi

echo "Getting python command in virual environment"
if [ -f $PYTHON_VENV_EXEC_PATH/python3 ]; then
    PYTHON_CMD_PATH=$PYTHON_VENV_EXEC_PATH/python3
elif [ -f $PYTHON_VENV_EXEC_PATH/python ]; then
    PYTHON_CMD_PATH=$PYTHON_VENV_EXEC_PATH/python
else
    echo "Error: python command not found" >&2
    exit 1
fi

echo "Found python command"

echo "Activating Python virtual environment"
source $PYTHON_VENV_EXEC_PATH/activate

echo "Getting pip command"
PIP_CMD=pip3
$PIP_CMD --version > /dev/null 2>&1
if [ ! $? -eq 0 ]; then
    PIP_CMD=pip
fi

$PIP_CMD --version > /dev/null 2>&1
if [ ! $? -eq 0 ]; then
    echo "Error: pip command not found, please install Python package manager first" >&2
    exit 1
fi

echo "Found pip command"
$PIP_CMD --version

echo "Verifying Python virtual environment creation"
if [ -z $VIRTUAL_ENV ]; then
    echo "Error: Python virtual environment setup failed" >&2
    exit 1
fi

echo "Installing Python dependencies"
$PIP_CMD install -r $SCRIPT_DIR/gui/requirements.txt

echo "Installed Python dependencies"
$PIP_CMD freeze

echo "Getting go command"
go version
if [ ! $? -eq 0 ]; then
    echo "Go command not found, please install Go first" >&2
    exit 1
fi

echo "Building Go CLI"
cd $SCRIPT_DIR/cli/ && go build && cd -

echo "Go CLI build completed"

echo "Successfully completed setup"
