let scene, camera, renderer, controls;
let cubes = [];

async function init() {
  // Scene
  scene = new THREE.Scene();
  scene.background = new THREE.Color(0x333333);

  // Camera
  camera = new THREE.PerspectiveCamera(
    75,
    window.innerWidth / window.innerHeight,
    0.1,
    1000,
  );
  camera.position.set(15, 15, 15);

  // Renderer
  renderer = new THREE.WebGLRenderer({ antialias: true });
  renderer.setSize(window.innerWidth, window.innerHeight);
  document.body.appendChild(renderer.domElement);

  // Controls
  controls = new THREE.OrbitControls(camera, renderer.domElement);
  controls.enableDamping = true;
  controls.dampingFactor = 0.05;

  const ambientLight = new THREE.AmbientLight(0xffffff, 0.6);
  scene.add(ambientLight);

  const directionalLight = new THREE.DirectionalLight(0xffffff, 0.8);
  directionalLight.position.set(10, 20, 10);
  scene.add(directionalLight);

  console.log("YEEET");
  // Grid helper
  const gridHelper = new THREE.GridHelper(20, 20);
  scene.add(gridHelper);

  try {
    // Load and display data
    const response = await fetch("/api/cubes");
    if (!response.ok) {
      throw new Error("Failed to fetch data");
    }
    const data = await response.json();

    // Container outline
    const containerGeometry = new THREE.BoxGeometry(
      data.Width,
      data.Height,
      data.Depth,
    );
    const edges = new THREE.EdgesGeometry(containerGeometry);
    const containerOutline = new THREE.LineSegments(
      edges,
      new THREE.LineBasicMaterial({ color: 0xffffff }),
    );
    containerOutline.position.set(
      data.Width / 2,
      data.Height / 2,
      data.Depth / 2,
    );
    scene.add(containerOutline);

    // Cubes
    data.Cubes.forEach((cubeData, index) => {
      const geometry = new THREE.BoxGeometry(
        cubeData.Width,
        cubeData.Height,
        cubeData.Depth,
      );
      const material = new THREE.MeshPhongMaterial({
        color: new THREE.Color().setHSL(index * 0.1, 0.7, 0.5),
        transparent: true,
        opacity: 0.8,
      });
      const cube = new THREE.Mesh(geometry, material);

      // Position cube
      cube.position.set(
        cubeData.X + cubeData.Width / 2,
        cubeData.Y + cubeData.Height / 2,
        cubeData.Z + cubeData.Depth / 2,
      );

      scene.add(cube);
      cubes.push(cube);
    });

    // GUI controls
    const gui = new dat.GUI();
    const params = {
      opacity: 0.8,
      wireframe: false,
    };

    gui.add(params, "opacity", 0, 1).onChange((value) => {
      cubes.forEach((cube) => {
        cube.material.opacity = value;
      });
    });

    gui.add(params, "wireframe").onChange((value) => {
      cubes.forEach((cube) => {
        cube.material.wireframe = value;
      });
    });
  } catch (error) {
    console.error("Error loading data:", error);
    document.body.innerHTML = `<div style="color: white; padding: 20px;">Error loading data: ${error.message}</div>`;
  }

  animate();
}

function animate() {
  requestAnimationFrame(animate);
  controls.update();
  renderer.render(scene, camera);
}

function onWindowResize() {
  camera.aspect = window.innerWidth / window.innerHeight;
  camera.updateProjectionMatrix();
  renderer.setSize(window.innerWidth, window.innerHeight);
}

window.addEventListener("resize", onWindowResize, false);
window.addEventListener("load", init);

// error handling
window.addEventListener("error", function (e) {
  if (e.target.tagName === "SCRIPT") {
    document.body.innerHTML = `
            <div style="color: white; padding: 20px;">
                Error loading required scripts.
            </div>
        `;
  }
});
