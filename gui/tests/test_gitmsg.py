import sys
from unittest import TestCase, mock
import pyperclip
from PyQt6.QtWidgets import QApplication
from PyQt6.QtTest import QTest
from PyQt6.QtCore import Qt
from gitmsg import GitmsgGUI


app = QApplication(sys.argv)


class TestGitmsgGUI(TestCase):
    def setUp(self):
        self.window = GitmsgGUI()

    def test_msg(self):
        summary = 'This is a summary'
        body = [
            'This is an incredibly detailed, unnecessarily long, pretentious,',
            'obnoxious paragraph',
        ]
        expected_msg = summary + '\n\n' + '\n'.join(body)
        self.window.summary.setText(summary)
        self.window.body.setText(' '.join(body))

        self.assertEqual(self.window.msg, expected_msg)

    def test_export_msg(self):
        summary = 'This is a summary'
        body = [
            'This is an incredibly detailed, unnecessarily long, pretentious,',
            'obnoxious paragraph',
        ]
        expected_msg = summary + '\n\n' + '\n'.join(body)
        self.window.summary.setText(summary)
        self.window.body.setText(' '.join(body))

        m = mock.mock_open()
        with mock.patch('builtins.open', m):
            QTest.mouseClick(self.window.export_button, Qt.MouseButton.LeftButton)

        m.assert_called_once_with(self.window.export_file_name, 'w')
        handle = m()
        handle.write.assert_called_once_with(expected_msg)

    def test_copy_msg(self):
        summary = 'This is a summary'
        body = [
            'This is an incredibly detailed, unnecessarily long, pretentious,',
            'obnoxious paragraph',
        ]
        expected_msg = summary + '\n\n' + '\n'.join(body)
        self.window.summary.setText(summary)
        self.window.body.setText(' '.join(body))

        m = mock.MagicMock()
        with mock.patch('pyperclip.copy', m):
            QTest.mouseClick(self.window.copy_button, Qt.MouseButton.LeftButton)

        m.assert_called_once_with(expected_msg)
