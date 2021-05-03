package playermove

type Player struct {
	currPos Vec2
	targetPos Vec2
	speed float32
}

//设置玩家目标位置
func (p *Player) MoveTo(v Vec2)  {
	p.targetPos = v
}

//获取当前位置
func (p Player) Pos() Vec2 {
	return p.currPos
}

//判断是否到达目的地
func (p *Player) IsArrived() bool {
	//通过计算当前玩家位置与目标位置距离不超过移动的步长，判断已经到达终点
	return p.currPos.DistanceTo(p.targetPos) < p.speed
}

//更新玩家位置
func (p *Player) Update() {
	if !p.IsArrived() {
		//计算出当前位置指向目标的朝向
		dir := p.targetPos.Sub(p.currPos).Normalize()
		//添加速度矢量生成新的位置
		newPos := p.currPos.Add(dir.Scale(p.speed))
		//更新
		p.currPos = newPos
	}
}

//创建新玩家
func NewPlayer(speed float32) *Player {
	return &Player{
		speed: speed,
	}
}






