function drawMatrixRain() {
    const canvas = document.createElement('canvas');
    document.body.appendChild(canvas);
    const ctx = canvas.getContext('2d');

    canvas.width = window.innerWidth;
    canvas.height = window.innerHeight;

    const columns = Array.from({ length: Math.floor(canvas.width / 20) }, () => 0);

    function draw() {
        ctx.fillStyle = 'rgba(0, 0, 0, 0.05)';
        ctx.fillRect(0, 0, canvas.width, canvas.height);

        ctx.fillStyle = '#00ff00';
        ctx.font = '20px Courier';

        columns.forEach((y, index) => {
            const text = String.fromCharCode(Math.floor(Math.random() * 128));
            const x = index * 20;
            ctx.fillText(text, x, y);
            columns[index] = y >= canvas.height || Math.random() > 0.95 ? 0 : y + 20;
        });
    }

    setInterval(draw, 50);
}

drawMatrixRain();
