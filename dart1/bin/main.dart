import 'dart:io';

void main() {
  stdout.writeln(Platform.environment);
  
  String input = stdin.readLineSync();
  stdout.writeln('You typed: $input');
}