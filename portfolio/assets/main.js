document.addEventListener("DOMContentLoaded", (event) => {
    var canvas = document.getElementById("decision-cube");
    if (!canvas) {
        console.log("Couldn't retrieve canvas.");
        return;
    }

    var scene = new THREE.Scene();
    scene.background = new THREE.Color( 0xffffff ); 
    var camera = new THREE.PerspectiveCamera( 75, window.innerWidth/window.innerHeight, 0.1, 1000 );

    var renderer = new THREE.WebGLRenderer({
        'canvas': document.getElementById("decision-cube"),
    });

    //renderer.setSize( window.innerWidth, window.innerHeight );

    var geometry = new THREE.BoxGeometry(1, 1, 1);
    //var texture = new THREE.TextureLoader().load('assets/logo.svg');
    var material = new THREE.MeshBasicMaterial({
        color: 0x111111,
        wireframe: true,
        //map: texture,
    });
    var cube = new THREE.Mesh( geometry, material );
    scene.add( cube );

    camera.position.z = 5;

    var animate = function () {
        requestAnimationFrame( animate );

        cube.rotation.x += 0.01;
        cube.rotation.y += 0.01;

        renderer.render(scene, camera);
    };

    animate();
    console.log("going");
});

