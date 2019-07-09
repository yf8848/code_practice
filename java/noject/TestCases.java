import java.util.HashMap;
import java.util.Map;

public class TestCases{


	public static void main(String[] args) {
		TestForEachMap();
	}

	public static void TestForEachMap(){
		Map<String, String> forMap = new HashMap<String, String>();
		forMap.put("one","zhangsan");
		forMap.put("two","lisi");
		forMap.put("three", "wangwu");

		forMap.values().forEach(v->{
			System.out.println(v);
		});
	}
}
