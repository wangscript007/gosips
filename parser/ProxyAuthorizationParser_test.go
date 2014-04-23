package parser

import (
	"testing"
)

func TestProxyAuthorizationParser(t *testing.T) {
	var inputs = []string{
		"Proxy-Authorization: Digest realm=\"MCI WorldCom SIP\"," +
			"domain=\"sip:ss2.wcom.com\", nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\"," +
			"opaque=\"\", stale=FALSE, algorithm=MD5\n",

		"Proxy-Authenticate: Digest realm=\"MCI WorldCom SIP\"," +
			"qop=\"auth\" , nonce-value=\"oli\"\n",
	}

	for i := 0; i < len(inputs); i++ {
		shp := NewProxyAuthorizationParser(inputs[i])
		testHeaderParser(t, shp)
	}
}

/** Test program
        public static void main(String args[]) throws ParseException {
		String paAuth[] = {
     "Proxy-Authorization: Digest realm=\"MCI WorldCom SIP\","+
     "domain=\"sip:ss2.wcom.com\", nonce=\"ea9c8e88df84f1cec4341ae6cbe5a359\","+
     "opaque=\"\", stale=FALSE, algorithm=MD5\n",

     "Proxy-Authenticate: Digest realm=\"MCI WorldCom SIP\","+
	"qop=\"auth\" , nonce-value=\"oli\"\n"
                };

		for (int i = 0; i < paAuth.length; i++ ) {
		    ProxyAuthorizationParser pap =
			  new ProxyAuthorizationParser(paAuth[i]);
		    ProxyAuthorization pa= (ProxyAuthorization) pap.parse();
		    System.out.println("encoded = " + pa.encode());
		}

	}
*/