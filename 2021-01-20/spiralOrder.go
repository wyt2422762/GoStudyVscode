package maximumproduct

//54. 螺旋矩阵

// 给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。

// 示例 1：
// 输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
// 输出：[1,2,3,6,9,8,7,4,5]

// 示例 2：
// 输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// 输出：[1,2,3,4,8,12,11,10,9,5,6,7]

func spiralOrder(matrix [][]int) []int {
    hang := len(matrix) //行数
    lie := len(matrix[0]) //列数

    res := make([]int,  hang * lie) //返回结果

    h, l := 0, 0 //起始行、列
    direction := "right" //方向，初始向右
    circle := 0 //第几轮

    for i := 0; i < hang * lie; i++{
        res[i] = matrix[h][l]

        //处理方向变化，下方的边界和轮次有关
        //向右到达右边界，方向改为向下
        if "right" == direction && l == lie - circle - 1 {
            direction = "down"
        }
        //向下到达下边界，方向改为向左
        if "down" == direction && h == hang - circle - 1 {
            direction = "left"
        }
        //向左到达左边界，方向改为线上
        if "left" == direction && l == 0 + circle {
            direction = "up"
        }
        //向上到达上边界
        //方向向上，且行数的为上一轮的起始行数+1，说明已经走完一圈，这里开启下一圈(内圈)的处理
        if "up" == direction && h == 0 + circle + 1 {
            l++
            circle++
            direction = "right"
            continue
        }

        //行进
        if "right" == direction {
            l++
        } 
        if "down" == direction {
            h++
        }
        if "left" == direction {
            l--
        }
        if "up" == direction {
            h--
        }
    }
    return res
}