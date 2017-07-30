package main

func managePlayerFloorCollision(p *player, f *floor) {
	p.mu.Lock()
	f.mu.Lock()
	defer p.mu.Unlock()
	defer f.mu.Unlock()
	if p.x-p.w/2 < float32(f.wall) {
		p.dx = 0
		p.x = float32(f.wall) + p.w/2
		return
	}
	if p.x+p.w/2 > float32(f.w-f.wall) {
		p.dx = 0
		p.x = float32(f.w-f.wall) - p.w/2
		return
	}
	if p.y < p.w/2 {
		p.dy = 0
		p.y = p.w / 2
		return
	}
	if p.y > float32(f.h)-p.w/2 {
		p.dy = 0
		p.y = float32(f.h) - p.w/2
		return
	}

}
