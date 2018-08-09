class LinkedList
  def initialize(head = nil)
    @head = head
    if head
      @count = 1;
    else
      @count = 0;
    end
  end

  def getCount
    @count
  end

  def print_list
    _current = @head
    unless _current
      puts "List is empty!"
    else
      while _current do
        puts "Node: #{_current.getValue}"
        _current = _current.getNext
      end
    end
  end

  def insert(node)
    unless @head
      @head = node
      @count += 1
    else
      _current = @head
      while _current.getNext do
        _current = _current.getNext
      end
      _current.setNext(node)
      @count += 1
    end
  end

  def insert_beginning(node)
    unless @head
      @head = node
      @count += 1
    else
      node.setNext(@head)
      @head = node
      @count += 1
    end
  end

  def insert_after(node, afterNode)
    _current = @head
    if _current.getValue == afterNode.getValue
      node.setNext(_current.getNext)
      _current.setNext(node)
      @count += 1
    else
      while _current && _current.getNext do
        if _current.getNext.getValue == afterNode.getValue
          node.setNext(_current.getNext.getNext)
          _current.getNext.setNext(node)
          @count += 1
          return
        else
          _current = _current.getNext
        end
      end
    end
  end

  def remove(node)
    if @head && @head.getValue == node.getValue
      @head = @head.getNext
      @count -= 1
    else
      _current = @head
      while _current && _current.getNext do
        if _current.getNext.getValue == node.getValue
          _current.setNext(_current.getNext.getNext)
          @count -= 1
        else
          _current = _current.getNext
        end
      end
    end
  end

  def pop
    _current = @head
    if _current && !_current.getNext
      @head = nil
      @count = 0
    else
      while _current && _current.getNext do
        unless _current.getNext.getNext
          _current.setNext(nil)
          @count -= 1
          return
        else
          _current = _current.getNext
        end
      end
    end
  end

end
