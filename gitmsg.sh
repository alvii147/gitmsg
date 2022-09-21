SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

PYTHON_VENV_NAME=.venv
<<<<<<< HEAD
PYTHON_VENV_PATH=$SCRIPT_DIR/gui/$PYTHON_VENV_NAME
=======
PYTHON_VENV_PATH=$SCRIPT_DIR/gitmsgGUI/$PYTHON_VENV_NAME
>>>>>>> 1eddef0f2c75bc1c9cc7cd16349773e9f0bae73e

if [ -d $PYTHON_VENV_PATH/bin ]; then
    PYTHON_VENV_EXEC_PATH=$PYTHON_VENV_PATH/bin
elif [ -d $PYTHON_VENV_PATH/Scripts ]; then
    PYTHON_VENV_EXEC_PATH=$PYTHON_VENV_PATH/Scripts
else
    echo "Error: Python venv binaries not found, please run setup.sh" >&2
    exit 1
fi

if [ -f $PYTHON_VENV_EXEC_PATH/python3 ]; then
    PYTHON_CMD_PATH=$PYTHON_VENV_EXEC_PATH/python3
elif [ -f $PYTHON_VENV_EXEC_PATH/python ]; then
    PYTHON_CMD_PATH=$PYTHON_VENV_EXEC_PATH/python
else
    echo "Error: python command not found, please run setup.sh" >&2
    exit 1
fi

if [ -f $SCRIPT_DIR/cli/gitmsg ]; then
    GO_EXEC_PATH=$SCRIPT_DIR/cli/gitmsg
elif [ -f $SCRIPT_DIR/cli/gitmsg.exe ]; then
    GO_EXEC_PATH=$SCRIPT_DIR/cli/gitmsg.exe
else
    echo "Golang executable not found, please run setup.sh" >&2
    exit 1
fi

function print_help() {
    echo -e "gitmsg helps write git commit messages based on the git 50/72 rule."
    echo -e "\nUsage:"
    echo -e "\tgitmsg [OPTIONS]"
    echo -e "\nOptions:"
    echo -e "\t-g, run in GUI mode"
    echo -e "\t-c, run in CLI mode"
    echo -e "\t-h, show help"
}

GUI_MODE=0
CLI_MODE=0

while getopts "gch" flag; do
    case "$flag" in
        g)  GUI_MODE=1;;
        c)  CLI_MODE=1;;
        h)  print_help
            exit 0;;
        *)  exit 128;;
    esac
done

if [[ ! $((GUI_MODE + CLI_MODE)) -eq 1 ]]; then
    echo "Error: please choose a mode: -g (GUI mode) or -c (CLI mode)"
    print_help
    exit 1
fi

if [[ "$GUI_MODE" -eq 1 ]]; then
    source $PYTHON_VENV_EXEC_PATH/activate
    $PYTHON_CMD_PATH $SCRIPT_DIR/gui/gitmsg.py
elif [[ "$CLI_MODE" -eq 1 ]]; then
    $GO_EXEC_PATH
else
    echo "Error: please choose a mode: -g (GUI mode) or -c (CLI mode)"
    print_help
    exit 1
fi
