package github.why168.gomobile

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.Toast
import github.Github
import kotlin.concurrent.thread

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
    }

    /**
     * 启动webService之后很卡，但是不影响访问
     * 本机浏览器输入：http://localhost:5555/、http://localhost:5555/hello
     */
    fun startWeb(view: View) {
        val msg = "Github.startWeb()"
        Github.startWeb()
        Log.d("GoLang", msg)
        Toast.makeText(this, msg, Toast.LENGTH_SHORT).show()
    }

    fun add(view: View) {
        val msg = "Github.add(1, 4) = ${Github.add(1, 4)}"
        Log.d("GoLang", msg)
        Toast.makeText(this, msg, Toast.LENGTH_SHORT).show()

    }

    fun md5(view: View) {
        val msg = "Github.md5(\"1234\") = ${Github.md5("1234")}"
        Log.d("GoLang", msg)
        Toast.makeText(this, msg, Toast.LENGTH_SHORT).show()
    }
}
