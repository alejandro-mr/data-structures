class Node
  def initialize(value, nextNode = nil)
    @value = value
    @next = nextNode
  end

  def getValue
    @value
  end

  def getNext
    @next
  end

  def setValue(value)
    @value = value
  end

  def setNext(node)
    @next = node
  end
end
