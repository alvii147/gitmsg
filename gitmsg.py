import os
import sys
import subprocess

from PyQt5.QtWidgets import (
    QMainWindow,
    QApplication,
    QWidget,
    QVBoxLayout,
    QHBoxLayout,
    QLabel,
    QLineEdit,
    QTextEdit,
)
from PyQt5.QtCore import (
    Qt,
    QEasingCurve,
)
from PyQt5.QtGui import (
    QPixmap,
    QIcon,
)

from MangoUI import Button, Slider

import textwrap

class Window(QMainWindow):
    def __init__(self):
        super().__init__()
        self.initConfig()
        self.initStyles()
        self.initTextWrapper()
        self.initUI()

    def initConfig(self):
        self._width = 1400
        self._height = 800
        self._xPos = 330
        self._yPos = -900

        self.scriptDir = os.path.dirname(os.path.realpath(__file__))
        self.imgLogoSmallPath = os.path.join(self.scriptDir, 'img/gitmsg_logo_small.png')
        self.imgIconPath = os.path.join(self.scriptDir, 'img/gitmsg_icon.png')

        self.summaryLimit = 50
        self.bodyWrapLimit = 72

        self.exportFilename = 'gitmsg.txt'

    def initStyles(self):
        self.primaryBackgroundColor = 'rgb(23, 11, 59)'
        self.secondaryBackgroundColor = 'rgb(52, 25, 72)'
        self.primaryColor = 'rgb(13, 0, 26)'
        self.secondaryColor = 'rgb(255, 51, 153)'
        self.textColor = 'rgb(255, 255, 255)'

        self.fontFamily = 'Consolas'
        self.fontSize = 11

        self.borderColor = 'rgb(128, 0, 64)'
        self.borderRadius = 3

    def initTextWrapper(self):
        self.msg = ''
        self.wrapper = textwrap.TextWrapper(width = self.bodyWrapLimit, replace_whitespace = False)

    def initUI(self):
        self.setGeometry(self._xPos, self._yPos, self._width, self._height)
        self.setWindowTitle('gitmsg')
        self.setWindowIcon(QIcon(self.imgIconPath))

        self.setStyleSheet(f'''
            QMainWindow {{
                background: QLinearGradient(x1:0 y1:0, x2:1 y2:0, stop:0 {self.primaryBackgroundColor}, stop:1 {self.secondaryBackgroundColor});
            }}
        ''')

        self.mainLayout = QVBoxLayout()
        self.editorLayout = QHBoxLayout()

        self.inputsLayout = QVBoxLayout()
        self.actionButtonsLayout = QHBoxLayout()

        self.previewLayout = QVBoxLayout()

        self.summaryLabel = QLabel()
        self.summaryLabel.setText('Summary')
        self.summaryLabel.setStyleSheet(f'''
            QLabel {{
                color: {self.textColor};
                font-size: {self.fontSize}pt;
            }}
        ''')

        self.summary = QLineEdit()
        self.summary.setStyleSheet(f'''
            QLineEdit {{
                color: {self.textColor};
                background-color: {self.primaryColor};
                font-size: {self.fontSize}pt;
                border: 1px solid {self.borderColor};
                border-radius: {self.borderRadius}px;
            }}
        ''')
        self.summary.setMaxLength(self.summaryLimit)
        self.summary.textChanged.connect(self.displayMsg)

        self.bodyLabel = QLabel()
        self.bodyLabel.setText('Body')
        self.bodyLabel.setStyleSheet(f'''
            QLabel {{
                color: {self.textColor};
                font-size: {self.fontSize}pt;
            }}
        ''')

        self.body = QTextEdit()
        self.body.setStyleSheet(f'''
            QTextEdit {{
                color: {self.textColor};
                background-color: {self.primaryColor};
                font-size: {self.fontSize}pt;
                border: 1px solid {self.borderColor};
                border-radius: {self.borderRadius}px;
            }}
        ''')
        self.body.textChanged.connect(self.displayMsg)

        self.exportButton = Button(
            primaryColor = self.secondaryColor,
            secondaryColor = self.primaryBackgroundColor,
            parentBackgroundColor = self.primaryColor,
            borderWidth = 1,
            borderRadius = 3,
            fontSize = self.fontSize,
        )
        self.exportButton.setText('Export')
        self.exportButton.clicked.connect(self.exportMsg)

        self.commitButton = Button(
            primaryColor = self.secondaryColor,
            secondaryColor = self.primaryBackgroundColor,
            parentBackgroundColor = self.primaryColor,
            borderWidth = 1,
            borderRadius = 3,
            fontSize = self.fontSize,
        )
        self.commitButton.setText('Commit')
        self.commitButton.clicked.connect(self.commitMsg)

        self.logoPixmap = QPixmap(self.imgLogoSmallPath)
        self.logoPixmap.scaledToWidth(20)

        self.logoLabel = QLabel()
        self.logoLabel.setPixmap(self.logoPixmap)

        self.previewLabel = QLabel()
        self.previewLabel.setText('Preview')
        self.previewLabel.setStyleSheet(f'''
            QLabel {{
                color: {self.textColor};
                font-size: {self.fontSize}pt;
            }}
        ''')

        self.preview = QTextEdit()
        self.preview.setStyleSheet(f'''
            QTextEdit {{
                color: {self.secondaryColor};
                background-color: {self.primaryColor};
                font-family: {self.fontFamily};
                font-size: {self.fontSize}pt;
                border: 1px solid {self.borderColor};
                border-radius: {self.borderRadius}px;
            }}
        ''')
        self.preview.setReadOnly(True)

        self.inputsLayout.addWidget(self.summaryLabel)
        self.inputsLayout.addWidget(self.summary)

        self.inputsLayout.addWidget(self.bodyLabel)
        self.inputsLayout.addWidget(self.body)

        self.previewLayout.addWidget(self.previewLabel)
        self.previewLayout.addWidget(self.preview)

        self.editorLayout.addLayout(self.inputsLayout)
        self.editorLayout.addLayout(self.previewLayout)

        self.actionButtonsLayout.addWidget(self.exportButton)
        self.actionButtonsLayout.addWidget(self.commitButton)
        self.actionButtonsLayout.addStretch()
        self.actionButtonsLayout.addWidget(self.logoLabel)

        self.mainLayout.addLayout(self.editorLayout)
        self.mainLayout.addLayout(self.actionButtonsLayout)

        self.centralWidget = QWidget(self)
        self.centralWidget.setLayout(self.mainLayout)
        self.setCentralWidget(self.centralWidget)

        self.show()

    def displayMsg(self):
        summary = self.summary.text()[:self.summaryLimit]
        if len(summary.strip()) > 0:
            summary += '\n\n'

        body = self.body.toPlainText()
        body = ['\n'.join(self.wrapper.wrap(text = block)) for block in body.split('\n')]
        body = '\n'.join(body)

        self.msg = summary + body
        self.preview.setText(self.msg)

    def exportMsg(self):
        self.displayMsg()
        with open(self.exportFilename, 'w') as exportFile:
            exportFile.write(self.msg)

    def commitMsg(self):
        self.exportMsg()
        cmd = f'git commit -F {self.exportFilename}'
        subprocess.run(cmd.split())

if __name__ == '__main__':
    app = QApplication(sys.argv)
    myWin = Window()
    sys.exit(app.exec_())