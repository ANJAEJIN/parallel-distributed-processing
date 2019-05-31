class MergeSortAlgorithm
  def merge_sort(arr)
    return arr unless arr.size > 1
    mid = arr.size/2
    a, b, sorted = merge_sort(arr[0...mid]), merge_sort(arr[mid..-1]), []
    sorted << (a[0] < b[0] ? a.shift : b.shift) while [a,b].none?(&:empty?)
    sorted + a + b
  end
end

merge_sort = MergeSortAlgorithm.new

arr = IO.readlines("random.txt")
int_array = arr.map(&:to_i)

start = Time.now
puts merge_sort.merge_sort(int_array)
finish = Time.now
puts(finish - start)