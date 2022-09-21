import os
import sys
import subprocess

from PyQt6.QtWidgets import (
    QMainWindow,
    QApplication,
    QWidget,
    QVBoxLayout,
    QHBoxLayout,
    QLabel,
    QLineEdit,
    QTextEdit,
)
from PyQt6.QtCore import (
    Qt,
    QEasingCurve,
)
from PyQt6.QtGui import (
    QPixmap,
    QIcon,
    QPalette,
    QColor,
)

from MangoUI import Button, Slider

import textwrap

class Window(QMainWindow):
    def __init__(self):
        super().__init__()
        self.init_config()
        self.init_styles()
        self.init_text_wrapper()
        self.init_ui()

    def init_config(self):
        self._width = 800
        self._height = 600
        self._xPos = 330
        self._yPos = 200

        self.script_dir = os.path.dirname(os.path.realpath(__file__))
        self.img_logo_small_path = os.path.join(self.script_dir, '../img/gitmsg_logo_small.png')
        self.img_icon_path = os.path.join(self.script_dir, '../img/gitmsg_icon.png')

        self.summary_limit = 50
        self.body_wrap_limit = 72

        self.export_file_name = 'gitmsg.txt'

    def init_styles(self):
        self.primary_background_color = 'rgb(23, 11, 59)'
        self.secondary_background_color = 'rgb(52, 25, 72)'
        self.primary_color = 'rgb(13, 0, 26)'
        self.secondary_color = 'rgb(255, 51, 153)'
        self.text_color = 'rgb(255, 255, 255)'

        self.font_family = 'Consolas'
        self.font_size = 11

        self.border_color = 'rgb(128, 0, 64)'
        self.border_radius = 3

    def init_text_wrapper(self):
        self.msg = ''
        self.wrapper = textwrap.TextWrapper(width = self.body_wrap_limit, replace_whitespace = False)

    def init_ui(self):
        self.setGeometry(self._xPos, self._yPos, self._width, self._height)
        self.setWindowTitle('gitmsg')
        self.setWindowIcon(QIcon(self.img_icon_path))

        self.setStyleSheet(f'''
            QMainWindow {{
                background: QLinearGradient(x1:0 y1:0, x2:1 y2:0, stop:0 {self.primary_background_color}, stop:1 {self.secondary_background_color});
            }}
        ''')

        self.main_layout = QVBoxLayout()
        self.editor_layout = QHBoxLayout()

        self.inputs_layout = QVBoxLayout()
        self.action_buttons_layout = QHBoxLayout()

        self.preview_layout = QVBoxLayout()

        self.summary_label = QLabel()
        self.summary_label.setText('Summary')
        self.summary_label.setStyleSheet(f'''
            QLabel {{
                color: {self.text_color};
                font-size: {self.font_size}pt;
            }}
        ''')

        self.summary = QLineEdit()
        self.summary.setStyleSheet(f'''
            QLineEdit {{
                color: {self.text_color};
                background-color: {self.primary_color};
                font-size: {self.font_size}pt;
                border: 1px solid {self.border_color};
                border-radius: {self.border_radius}px;
            }}
        ''')
        self.summary.setMaxLength(self.summary_limit)
        self.summary.textChanged.connect(self.display_msg)

        self.body_label = QLabel()
        self.body_label.setText('Body')
        self.body_label.setStyleSheet(f'''
            QLabel {{
                color: {self.text_color};
                font-size: {self.font_size}pt;
            }}
        ''')

        self.body = QTextEdit()
        self.body.setStyleSheet(f'''
            QTextEdit {{
                color: {self.text_color};
                background-color: {self.primary_color};
                font-size: {self.font_size}pt;
                border: 1px solid {self.border_color};
                border-radius: {self.border_radius}px;
            }}
        ''')
        self.body.textChanged.connect(self.display_msg)

        self.export_button = Button(
            primaryColor = self.secondary_color,
            secondaryColor = self.primary_background_color,
            parentBackgroundColor = self.primary_color,
            borderWidth = 1,
            borderRadius = 3,
            fontSize = self.font_size,
        )
        self.export_button.setText('Export')
        self.export_button.clicked.connect(self.export_msg)

        self.commit_button = Button(
            primaryColor = self.secondary_color,
            secondaryColor = self.primary_background_color,
            parentBackgroundColor = self.primary_color,
            borderWidth = 1,
            borderRadius = 3,
            fontSize = self.font_size,
        )
        self.commit_button.setText('Commit')
        self.commit_button.clicked.connect(self.commit_msg)

        self.logo_pixmap = QPixmap(self.img_logo_small_path)
        self.logo_pixmap.scaledToWidth(20)

        self.logo_label = QLabel()
        self.logo_label.setPixmap(self.logo_pixmap)

        self.preview_label = QLabel()
        self.preview_label.setText('Preview')
        self.preview_label.setStyleSheet(f'''
            QLabel {{
                color: {self.text_color};
                font-size: {self.font_size}pt;
            }}
        ''')

        self.preview = QTextEdit()
        self.preview.setStyleSheet(f'''
            QTextEdit {{
                color: {self.secondary_color};
                background-color: {self.primary_color};
                font-family: {self.font_family};
                font-size: {self.font_size}pt;
                border: 1px solid {self.border_color};
                border-radius: {self.border_radius}px;
            }}
        ''')
        self.preview.setReadOnly(True)

        self.inputs_layout.addWidget(self.summary_label)
        self.inputs_layout.addWidget(self.summary)

        self.inputs_layout.addWidget(self.body_label)
        self.inputs_layout.addWidget(self.body)

        self.preview_layout.addWidget(self.preview_label)
        self.preview_layout.addWidget(self.preview)

        self.editor_layout.addLayout(self.inputs_layout)
        self.editor_layout.addLayout(self.preview_layout)

        self.action_buttons_layout.addWidget(self.export_button)
        self.action_buttons_layout.addWidget(self.commit_button)
        self.action_buttons_layout.addStretch()
        self.action_buttons_layout.addWidget(self.logo_label)

        self.main_layout.addLayout(self.editor_layout)
        self.main_layout.addLayout(self.action_buttons_layout)

        self.central_widget = QWidget(self)
        self.central_widget.setLayout(self.main_layout)
        self.setCentralWidget(self.central_widget)

        self.show()

    def display_msg(self):
        summary = self.summary.text()[:self.summary_limit]
        if len(summary.strip()) > 0:
            summary += '\n\n'

        body = self.body.toPlainText()
        body = ['\n'.join(self.wrapper.wrap(text = block)) for block in body.split('\n')]
        body = '\n'.join(body)

        self.msg = summary + body
        self.preview.setText(self.msg)

    def export_msg(self):
        self.display_msg()
        with open(self.export_file_name, 'w') as f:
            f.write(self.msg)

    def commit_msg(self):
        self.export_msg()
        cmd = f'git commit -F {self.export_file_name}'
        subprocess.run(cmd.split())

if __name__ == '__main__':
    app = QApplication(sys.argv)
    my_win = Window()
    sys.exit(app.exec())