#!/usr/bin/env ruby

require './LinkedList'
require './Node'

ll = LinkedList.new

ll.insert(Node.new(1))
ll.insert(Node.new(2))
ll.insert(Node.new(3))
ll.insert_after(Node.new(4), Node.new(3))
ll.print_list()
puts "Count: #{ll.getCount}\n\n"

ll.insert_beginning(Node.new(5))
ll.print_list()
puts "Count: #{ll.getCount}\n\n"

puts "About to remove node with value of 2"
ll.remove(Node.new(2))
ll.print_list()
puts "Count: #{ll.getCount}\n\n"

puts "Removing last element of list"
ll.pop()
ll.print_list()
puts "Count: #{ll.getCount}\n\n"

puts"Trying to remove non existen node with value 2"
ll.remove(Node.new(2))
ll.print_list()
puts "Count: #{ll.getCount}\n\n"

puts "About to empty list by calling `pop` multiple times"
(ll.getCount + 10).times {
  ll.pop()
}
ll.print_list()
puts "Count: #{ll.getCount}\n\n"
