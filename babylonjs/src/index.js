import {
    Scene,
    Engine,
    ArcRotateCamera,
    SceneLoader,
    Vector3,Color3,
    AnimationPropertiesOverride,
    ExecuteCodeAction,
    Mesh,
    ActionManager,
    PointLight,
    Tools
} from 'babylonjs';
import './styles.css'

let canvas = document.createElement("canvas")
canvas.classList.add("canvas")
document.body.appendChild(canvas)

let engine = new Engine(canvas, true)
engine.enableOfflineSupport = false;

let scene = new Scene(engine)
scene.clearColor = new Color3.White()

var plane = Mesh.CreatePlane("plane", 200, scene, false, Mesh.DOUBLESIDE);
plane.position = new Vector3(0,0,0);
plane.rotation = new Vector3(Tools.ToRadians(90), 0, 0)

let camera = new ArcRotateCamera("Cam1", Tools.ToRadians(90), Tools.ToRadians(90), 500, new Vector3(0,0,0), scene)
camera.attachControl(canvas,true);

var light = new PointLight("pointLight",new Vector3(0,100,-100),scene);
light.diffuse = new Color3(1,1,1);

SceneLoader.ImportMesh("", "./scenes/", "model.babylon", scene, (newMeshes, particleSystems, skeletons) => {
    let skeleton = skeletons[0]
    newMeshes.forEach(m=>m.rotation = new Vector3(Tools.ToRadians(-90), 0, 0))
    skeleton.animationPropertiesOverride = new AnimationPropertiesOverride()
    skeleton.animationPropertiesOverride.enableBlending = true
    skeleton.animationPropertiesOverride.blendingSpeed = 0.05
    skeleton.animationPropertiesOverride.loopMode = 1

    let runRange = skeleton.getAnimationRange("Run")

    if (runRange) scene.beginAnimation(skeleton, runRange.from, runRange.to, true);
})

scene.actionManager = new ActionManager(scene);
scene.actionManager.registerAction(new ExecuteCodeAction(
    {
        trigger: ActionManager.OnKeyUpTrigger,
        parameter: " ",
    },
    function () {
        light.setEnabled(!light.isEnabled());
    }
))

engine.runRenderLoop(()=>{
    scene.render()
})
