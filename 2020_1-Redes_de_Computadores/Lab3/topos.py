
from mininet.topo import Topo
from mininet.net import Mininet
from mininet.node import OVSSwitch, Controller, RemoteController

class Red1( Topo ):

    def __init__( self ):

        Topo.__init__( self )

        aHost = self.addHost( 'h1', mac= "00:00:00:00:00:01" )
        bHost = self.addHost( 'h2', mac= "00:00:00:00:00:02" )
        cHost = self.addHost( 'h3', mac= "00:00:00:00:00:03" )
        dHost = self.addHost( 'h4', mac= "00:00:00:00:00:04" )
        eHost = self.addHost( 'h5', mac= "00:00:00:00:00:05" )
        fHost = self.addHost( 'h6', mac= "00:00:00:00:00:06" )
        xSwitch = self.addSwitch( 's1' )
        ySwitch = self.addSwitch( 's2' )
        zSwitch = self.addSwitch( 's3' )
        """
        self.addLink( xSwitch, aHost )
        self.addLink( xSwitch, bHost )
        self.addLink( xSwitch, ySwitch )
        self.addLink( xSwitch, zSwitch )

        self.addLink( zSwitch, eHost )
        self.addLink( zSwitch, fHost )
        self.addLink( zSwitch, ySwitch)

        self.addLink( ySwitch, cHost )
        self.addLink( ySwitch, dHost )
        """


        self.addLink( aHost, xSwitch, 1, 2 )
        self.addLink( bHost, xSwitch, 3, 4 )
        self.addLink( ySwitch, xSwitch, 5, 6 )
        self.addLink( xSwitch, zSwitch, 7, 8 )

        self.addLink( eHost, zSwitch, 9, 10 )
        self.addLink( fHost, zSwitch, 11, 12 )
        self.addLink( zSwitch, ySwitch, 13, 14)

        self.addLink( cHost, ySwitch, 15, 16 )
        self.addLink( dHost, ySwitch, 17, 18 )


class Red2( Topo ):

    def __init__( self ):

        Topo.__init__( self )

        aHost = self.addHost( 'a', mac= "00:00:00:00:00:01" )
        bHost = self.addHost( 'b', mac= "00:00:00:00:00:02" )
        cHost = self.addHost( 'c', mac= "00:00:00:00:00:03" )
        dHost = self.addHost( 'd', mac= "00:00:00:00:00:04" )
        eHost = self.addHost( 'e', mac= "00:00:00:00:00:05" )
        fHost = self.addHost( 'f', mac= "00:00:00:00:00:06" )
        gHost = self.addHost( 'g', mac= "00:00:00:00:00:07" )
        tSwitch = self.addSwitch( 's1' )
        uSwitch = self.addSwitch( 's2' )
        vSwitch = self.addSwitch( 's3' )
        wSwitch = self.addSwitch( 's4' )
        xSwitch = self.addSwitch( 's5' )
        ySwitch = self.addSwitch( 's6' )
        zSwitch = self.addSwitch( 's7' )


        self.addLink( vSwitch, aHost, 1, 2 )
        self.addLink( vSwitch, bHost, 3, 4 )

        self.addLink( vSwitch, xSwitch, 5, 6 )
        self.addLink( vSwitch, uSwitch, 7, 8 )

        self.addLink( xSwitch, cHost, 9, 10 )
        self.addLink( xSwitch, dHost, 11, 12 )
        self.addLink( xSwitch, zSwitch, 13, 14 )
        self.addLink( xSwitch, wSwitch, 15, 16 )

        self.addLink( zSwitch, eHost, 17, 18 )
        self.addLink( zSwitch, fHost, 19, 20 )
        self.addLink( zSwitch, ySwitch, 21, 22 )
        
        self.addLink( ySwitch, wSwitch, 23, 24 )
        self.addLink( wSwitch, uSwitch, 25, 26 )
        self.addLink( uSwitch, tSwitch, 27, 28 )
        self.addLink( tSwitch, gHost, 29, 30 )
        self.addLink( tSwitch, vSwitch, 31, 32 )



topos = { 'Red1': ( lambda: Red1() ) , 'Red2': ( lambda: Red2() )}